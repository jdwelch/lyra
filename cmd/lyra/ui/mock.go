package ui

import (
	"fmt"
	"log"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mgutz/ansi"
)

// ShowSuccessMessage prints "Success!" plus optional message to STDOUT
// Might want to change that string globally later and this is
// easier than faffing with externalising strings properly
func ShowSuccessMessage(params ...string) {
	var msg = ""
	if len(params) > 0 {
		msg = params[0]
	}
	log.Println() //FIXME: (JD) This feels gross
	log.Println(ansi.Green+"▸ Success!"+ansi.Reset, msg)
	log.Println()
}

// ShowApplyStart Might want to unify the first thing that gets spit out after
// calling a command; this is sort of like the instruction text
// in a dialog box
func ShowApplyStart(field string, msg string) {
	log.Println()
	log.Println(ansi.Cyan+"▸ "+field+ansi.Reset, msg)
	log.Println()
}

// ProgressBar displays a simple text spinner-style
// progress indicator with a label. Fakes long-running operations
// to increase realism
func ProgressBar(label string, duration int, finalmsg bool) {
	// FIXME: (JD) This is totally unrealistic and inadequate
	s := spinner.New(spinner.CharSets[34], 100*time.Millisecond)
	if finalmsg {
		s.FinalMSG = label + "… Done.\n" // Newline is important!
	}
	s.Prefix = label + "… " // Leave a little space after the label
	s.Start()
	time.Sleep(time.Duration(duration) * time.Millisecond)
	s.Stop()
}

// ResourceSet is a human-friendly message
func ResourceSet(message string) {
	fmt.Println(ansi.Green+"[set resource] "+ansi.Reset, message)
}

// ResourceError is a human-friendly error message
func ResourceError(message string) {
	fmt.Println(ansi.LightRed+"[error] "+ansi.Reset, message)
}

// ResourceDestroy is a human-friendly message
func ResourceDestroy(message string) {
	fmt.Println(ansi.Magenta+"[destroy resource] "+ansi.Reset, message)
}

// Success prints "Success!" plus optional message
func Success(params ...string) {
	var msg = ""
	if len(params) > 0 {
		msg = params[0]
		fmt.Println(ansi.Green + "▸ Success! " + ansi.Reset + msg)
	}
}

// Failure tries to summarise what went wrong
// and provide guidance for recovery
func Failure(params ...string) {
	var msg = ""
	if len(params) > 0 {
		msg = params[0]
		fmt.Println(ansi.Red + "▸ Operation failed! " + ansi.Reset + msg)
	}
}

// Notice displays the result of a notice() function call
func Notice(message string) {
	fmt.Println(ansi.LightCyan + "[notice]\t" + ansi.Reset + message)
}

// NoopBanner is some visual sugar to indicate no-op operations
func NoopBanner() {
	fmt.Println(ansi.Blue + "==== DRY RUN MODE ====" + ansi.Reset)
}

// DescribeStep is the first thing that gets spit out after
// calling a command; this is sort of like the instruction text
// in a dialog box
func DescribeStep(msg string) {
	fmt.Println("\n" + ansi.Cyan + "▸ " + msg + ansi.Reset + "\n")
}

// DescribeStepWithField is the first thing that gets spit out after
// calling a command; this is sort of like the instruction text
// in a dialog box
func DescribeStepWithField(field string, msg string) {
	fmt.Println("\n" + ansi.Cyan + "▸ " + field + " " + ansi.Reset + msg + "\n")
}

// DiffAdd show that a thing will be added
func DiffAdd(message string) {
	fmt.Println(ansi.Green + "+ " + message + ansi.Reset)
}

// DiffRemove show that a thing will be removed
func DiffRemove(message string) {
	fmt.Println(ansi.Magenta + "- " + message + ansi.Reset)
}

// DiffChange show that a thing will be changed
func DiffChange(message string) {
	fmt.Println(ansi.Yellow + "~ " + message + ansi.Reset)
}

// DiffUnchanged shows that a thing will be not be changed
func DiffUnchanged(message string) {
	fmt.Println(ansi.LightBlack + "  " + message + ansi.Reset)
}

// DiffConflict shows a resource property where what's in the manifest/stored state doesn't match the running state
func DiffConflict(message string) {
	fmt.Println(ansi.LightRed + "! " + message + ansi.Reset)
}

// ExecOut passes through a line of stdout from an ssh-like thing
func ExecOut(target string, task string, message string) {
	fmt.Println(ansi.Green + "[" + target + "](" + task + "): " + ansi.Reset + message)
}

// ResourceSummary spits out a nice list about what happened
func ResourceSummary(added string, removed string, changed string, unchanged string) {
	fmt.Println("Resources " + ansi.Green + "added: " + added + ansi.Magenta + " removed: " + removed +
		ansi.Green + " changed: " + changed + ansi.LightBlack + " unchanged: " + unchanged + ansi.Reset)
}
