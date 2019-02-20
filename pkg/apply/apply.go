package apply

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/lyraproj/hiera/lookup"
	"github.com/lyraproj/hiera/provider"
	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/lyraproj/lyra/pkg/loader"
	"github.com/lyraproj/lyra/pkg/logger"
	"github.com/lyraproj/puppet-evaluator/eval"
	"github.com/lyraproj/puppet-evaluator/types"
	"github.com/lyraproj/servicesdk/serviceapi"
	"github.com/lyraproj/servicesdk/wfapi"
	"github.com/lyraproj/wfe/api"
	"github.com/lyraproj/wfe/service"
	"github.com/lyraproj/wfe/wfe"
	"gopkg.in/src-d/enry.v1"
	"io/ioutil"
	"os"
	"strings"

	// Ensure that lookup function properly loaded
	_ "github.com/lyraproj/hiera/functions"
)

// Applicator is used to apply workflows
type Applicator struct {
	HomeDir string
}

type cmdError string

func (e cmdError) Error() string {
	return string(e)
}

// ApplyWorkflowWithHieraData will apply the named workflow with the supplied hiera data
func (a *Applicator) ApplyWorkflowWithHieraData(workflowName string, hieraData map[string]string) {
	a.applyWithHieraData(workflowName, hieraData, wfapi.Upsert)
}

func (a *Applicator) applyWithHieraData(workflowName string, hieraData map[string]string, intent wfapi.Operation) {
	m := convertToDeepMap(hieraData)
	hclog.Default().Debug("converted map to hiera data", "m", m)
	v := eval.Wrap(nil, m).(eval.OrderedMap)
	tp := func(ic lookup.ProviderContext, key string, _ map[string]eval.Value) (eval.Value, bool) {
		return v.Get4(key)
	}
	lookup.DoWithParent(context.Background(), tp, nil, applyWithContext(workflowName, intent))
}

//convertToDeepMap converts a map[string]string with entries like {k:"aws.tags.created_by", v:"user@company.com"}
//to a recursive map[string]interface{} with a key created_by nested under a key tags nested under a key aws
//e.g. output map[aws:map[tags:map[created_by:person@company.com lifetime:2hrs] hello:hi]]
func convertToDeepMap(hieraData map[string]string) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range hieraData {
		current := output
		tokens := strings.Split(k, ".")
		len := len(tokens)
		for index, token := range tokens {
			if index == len-1 {
				current[token] = v
			} else {
				if _, ok := current[token]; !ok {
					current[token] = make(map[string]interface{})
				}
				current = current[token].(map[string]interface{})
			}
		}
	}
	return output
}

// DeleteWorkflowWithHieraData will delete the named workflow with the supplied hiera data
func (a *Applicator) DeleteWorkflowWithHieraData(workflowName string, hieraData map[string]string) {
	a.applyWithHieraData(workflowName, hieraData, wfapi.Delete)
}

// ApplyWorkflow will apply the named workflow getting hiera data from file
func (a *Applicator) ApplyWorkflow(workflowName, hieraDataFilename string, intent wfapi.Operation) (exitCode int) {
	if a.HomeDir != `` {
		if err := os.Chdir(a.HomeDir); err != nil {
			ui.Message("error", fmt.Errorf("Unable to change directory to '%s'", a.HomeDir))
			return 1
		}
	}

	defer func() {
		plugin.CleanupClients()
		logger.Get().Debug("all plugins cleaned up")
		if e := recover(); e != nil {
			if err, ok := e.(cmdError); ok {
				ui.Message("error", err)
				exitCode = 1
			} else {
				panic(e)
			}
		}
	}()

	lookupOptions := map[string]eval.Value{
		`path`:                      types.WrapString(hieraDataFilename),
		provider.LookupProvidersKey: types.WrapRuntime([]lookup.LookupKey{provider.Yaml, provider.Environment})}

	lookup.DoWithParent(context.Background(), provider.MuxLookup, lookupOptions, applyWithContext(workflowName, intent))
	return 0
}

func applyWithContext(workflowName string, intent wfapi.Operation) func(eval.Context) {
	return func(c eval.Context) {
		logger := logger.Get()
		loader := loader.New(logger, c.Loader())
		loader.PreLoad(c)
		logger.Debug("all plugins loaded")
		c.DoWithLoader(loader, func() {
			if intent == wfapi.Delete {
				logger.Debug("calling delete")
				delete(c, workflowName)
				ui.ShowMessage("delete done:", workflowName)
				logger.Debug("delete finished")
			} else {
				logger.Debug("calling apply")
				apply(c, workflowName, eval.EMPTY_MAP, intent) // TODO: Perhaps provide top-level input from command line args
				ui.ShowMessage("apply done:", workflowName)
				logger.Debug("apply finished")
			}
		})
	}
}

func loadActivity(c eval.Context, activityID string) api.Activity {
	def, ok := eval.Load(c, eval.NewTypedName(eval.NsDefinition, activityID))
	if !ok {
		panic(cmdError(fmt.Sprintf("Unable to find definition for activity %s", activityID)))
	}
	return wfe.CreateActivity(def.(serviceapi.Definition))
}

func delete(c eval.Context, activityID string) {
	log := logger.Get()
	log.Debug("deleting", "activityID", activityID)

	// Nothing in the workflow will be in the new era so all is deleted
	service.StartEra(c)
	service.SweepAndGC(c, loadActivity(c, activityID).Identifier()+"/")
}

func apply(c eval.Context, activityID string, input eval.OrderedMap, intent wfapi.Operation) {
	log := logger.Get()

	log.Debug("configuring scope")
	c.Scope().Set(service.ActivityContextKey, types.SingletonHash2(`operation`, types.WrapInteger(int64(intent))))

	log.Debug("applying", "activityID", activityID)
	service.StartEra(c)
	a := loadActivity(c, activityID)
	result := a.Run(c, eval.Wrap(c, input).(eval.OrderedMap))
	log.Debug("Apply done", "result", result)

	gcPrefix := a.Identifier() + "/"
	log.Debug("garbage collecting", "prefix", gcPrefix)
	service.SweepAndGC(c, gcPrefix)
}

func getLanguage(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return enry.GetLanguage(path, content), nil
}
