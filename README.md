# MongoDB Atlas Resource Provider for CloudFormation

This is the repository for the MongoDB Atlas Resource Provider for CloudFormation, which allows one to use CloudFormation to programmatically manage MongoDB Atlas, MongoDB's Database as a Service offering.
Learn more about Atlas at  [https://www.mongodb.com/cloud/atlas](https://www.mongodb.com/cloud/atlas)

For general information about CloudFormation, visit the [official website](https://aws.amazon.com/cloudformation) and the [GitHub project page](https://github.com/aws-cloudformation/).

# Support, Bugs, Feature Requests

Support for the MongoDB Atlas Resource Provider for CloudFormation is provided under MongoDB Atlas support plans, starting with Developer. Please submit support questions within the Atlas UI. Support questions submitted under the Issues section of this repo will be handled on a "best effort" basis.
Bugs should be filed under the Issues section of this repo.

Feature requests can be submitted at https://feedback.mongodb.com/forums/924145-atlas - just select the CloudFormation resource provider as the category or vote for an already suggested feature.

# Requirements
- [AWS CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli) 0.12+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

# Developing the CloudFormation Resource Provider

If you wish to work on the CloudFormation resource provider, you'll first need [Go](https://golang.org/doc/install) installed on your machine (please check the [requirements](#Requirements) before proceeding).

Note: This project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](https://golang.org/doc/code.html#GOPATH). The instructions that follow assume a directory in your home directory outside of the standard GOPATH (i.e $HOME/development/).

Clone repository to: `$HOME/development/`

```
$ cd $HOME/development/
$ git clone https://github.com/mongodb/mongodbatlas-cloudformation-resources.git
...
```

To develop a resource provider, you need to follow steps from [AWS CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli), you can read more info about developing a resource provider [here](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).

To create a resource, you need to create a folder first then generate files with `cfn init`

```
$ cd $HOME/development/mongodbatlas-cloudformation-resources
$ mkdir resource
$ cd resource
$ cfn init
...
```

# Using the CloudFormation Resource Provider

To use a resource provider, you need a template, templates of this project are available on `/examples`, then you need to run commands, you can read about available commands in [AWS CloudFormation commands](https://docs.aws.amazon.com/cli/latest/reference/cloudformation/index.html)

To create a stack of a resource provider

```
$ aws cloudformation create-stack --stack-name myTestProject --template-body file://../examples/project/project.json --parameters ParameterKey=Name,ParameterValue=test-project
```
If errors are not shown, it should return a stack id

To verify if it's working you can check on [AWS Cloudformation](https://console.aws.amazon.com/cloudformation)

# Testing the Provider

In order to test a provider resource, there are three ways to do that.

1. ### Using the AWS CloudFormation Console

   After creating, updating or deleting a stack of a provider resource in the AWS CloudFormation console, you can check the [AWS CloudFormation](https://console.aws.amazon.com/cloudformation) for events which will show the status of a resource provider.
      
   After uploading a stack logs are automatically created and may be viewed in [AWS CloudWatch](https://console.aws.amazon.com/cloudwatch) `Log groups`.
      
   To print a log you need to add it in code `log.Printf("")`.
   
2.  ### Using the AWS CloudFormation CLI
    
    After creating, updating or deleting a stack of a provider resource with the AWS CLI, you can use commands to see events, describe stacks and more. See [CloudFormation CLI commands](https://docs.aws.amazon.com/cli/latest/reference/cloudformation/index.html) for more.
        
    After uploading a stack logs are automatically created and may be viewed in [AWS CloudWatch](https://aws.amazon.com/cloudwatch/) `Log groups`.
        
    To retrieve logs you need the name of the `log group` and in some situations you need the name of the `log stream`. You can see more [here](https://aws.amazon.com/cloudwatch/) in `Log groups`, once there use [logs commads](https://docs.aws.amazon.com/cli/latest/reference/logs/index.html)

    
3.  ### Using the Contract Test

    The resources of this repository was created using CloudFormation best practices. Therefore, it will pass all of the CloudFormation CLI contract tests.
    To run the contrast test, follow the directions in the 
    [documentation](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-cli-test.html).
