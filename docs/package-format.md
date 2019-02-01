
# Package format

## Directory structure

```shell
my-project
├── README.md
├── metadata.yaml
├── deploy.yaml
├── values.yaml
└── workflows
    └──my-project-actions.ts
    └──my-project-resources.yaml
```

## Form and function of each file

**README.md**: A Markdown-formatted document for end-user consumption that includes installation and usage documentation, where to get help, and so on.

**metadata.yaml**: Properties of the workflow package.

_Example:_

```yaml
---
apiVersion: v1alpha1 #required?
appVersion: "0.1" #required?
name: My Workflow
author: jdwelch
description: "A lovely Lyra workflow."
version: 0.1.0 #required?
license: "Apache 2"
url: "gh.com/foo/bar"
```

**deploy.yaml**: A Kuberenetes object spec for deploying a workflow.

_Example:_

```yaml
---
apiVersion: lyraproj.io/v1alpha1
kind: Workflow
metadata:
  app.kubernetes.io/name: "my-workflow"
  labels:
    lyraproj.io/workflowVersion: 0.1.0
spec:
  lyraproj.io/workflowName: "my-workflow"
  lyraproj.io/data:
    aws_region: 'us-west-2'
    image_id: 'ami-b63ae0ce'
    refreshTime: 60
```

**values.yaml**: External data for the Lyra CLI to use as input when executing a workflow.

_Example:_

```yaml
---
aws_region: 'us-west-2'
image_id: 'ami-b63ae0ce'
```

**workflows/my-project-resources.ts**: Workflow manifest with sample action statement(s)

_Example:_

```typescript
[FIXME:]
```

**workflows/my-project-resources.yaml**: Workflow manifest with sample resource statement(s)

_Example:_

```yaml
[FIXME:]
```