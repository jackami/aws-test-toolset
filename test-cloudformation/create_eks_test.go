package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"io/ioutil"
	"testing"
)

func Test_CreateEKSNodeGroupStack(t *testing.T) {
	region := "us-west-2"
	stackName := "eksctl-eks-spot-eksctl-nodegroup-test"
	//templateName := "nodegroup_template.json"
	templateName := "nodegroup_spot_template.yaml"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	content, err := ioutil.ReadFile(templateName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Convert []byte to string
	//fmt.Println(string(content))
	templateBody := string(content)

	svc := cloudformation.New(sess)

	paras := []*cloudformation.Parameter{}

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("ClusterName"),
		ParameterValue: aws.String("eks-spot-eksctl"),
	})

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("ClusterControlPlaneSecurityGroup"),
		ParameterValue: aws.String("sg-id"),
	})

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("NodeGroupName"),
		ParameterValue: aws.String("test-test-aws-sdk-api-ng"),
	})

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("KeyName"),
		ParameterValue: aws.String("eks-spot-key"),
	})

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("VpcId"),
		ParameterValue: aws.String("vpc-id"),
	})

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("Subnets"),
		ParameterValue: aws.String("subnet-id-1,subnet-id-2,subnet-id-3"),
	})

	input := &cloudformation.CreateStackInput{
		Capabilities: aws.StringSlice([]string{"CAPABILITY_IAM"}),
		StackName:    aws.String(stackName),
		Parameters:   paras,
		TemplateBody: aws.String(templateBody),
	}

	//input = &cloudformation.CreateStackInput{
	//	Capabilities:                nil,
	//	ClientRequestToken:          nil,
	//	DisableRollback:             nil,
	//	EnableTerminationProtection: nil,
	//	NotificationARNs:            nil,
	//	OnFailure:                   nil,
	//	Parameters:                  nil,
	//	ResourceTypes:               nil,
	//	RoleARN:                     nil,
	//	RollbackConfiguration:       nil,
	//	StackName:                   nil,
	//	StackPolicyBody:             nil,
	//	StackPolicyURL:              nil,
	//	Tags:                        nil,
	//	TemplateBody:                nil,
	//	TemplateURL:                 nil,
	//	TimeoutInMinutes:            nil,
	//}

	result, err := svc.CreateStack(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

	fmt.Println("Waiting stack create complete ...")

	// snippet-start:[cfn.go.create_stack.wait]
	err = svc.WaitUntilStackCreateComplete(&cloudformation.DescribeStacksInput{
		StackName: aws.String(stackName),
	})
	if err != nil {
		fmt.Println("Got an error waiting for stack to be created")
		return
	}

	fmt.Println("finish create nodegroup")
}

func Test_DescribeStacks(t *testing.T) {
	region := "us-west-2"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := cloudformation.New(sess)
	input := &cloudformation.DescribeStacksInput{
		NextToken: nil,
		StackName: aws.String("eksctl-eks-spot-eksctl-nodegroup-dev-4vcpu-16gb-spot"),
	}

	result, err := svc.DescribeStacks(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func Test_GetTemplate(t *testing.T) {
	region := "us-west-2"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := cloudformation.New(sess)
	input := &cloudformation.GetTemplateInput{
		StackName: aws.String("eksctl-eks-spot-eksctl-nodegroup-dev-4vcpu-16gb-spot"),
	}

	result, err := svc.GetTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func Test_UpdateTemplateStack(t *testing.T) {
	region := "us-west-2"
	stackName := "eksctl-eks-spot-eksctl-nodegroup-test"
	//templateName := "nodegroup_template.json"
	templateName := "nodegroup_spot_template.yaml"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	content, err := ioutil.ReadFile(templateName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Convert []byte to string
	//fmt.Println(string(content))
	templateBody := string(content)

	svc := cloudformation.New(sess)

	paras := []*cloudformation.Parameter{}

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("ClusterName"),
		ParameterValue: aws.String("eks-spot-eksctl"),
	})

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("ClusterControlPlaneSecurityGroup"),
		ParameterValue: aws.String("sg-id"),
	})

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("NodeGroupName"),
		ParameterValue: aws.String("test-test-aws-sdk-api-ng"),
	})

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("KeyName"),
		ParameterValue: aws.String("eks-spot-key"),
	})

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("VpcId"),
		ParameterValue: aws.String("vpc-id"),
	})

	paras = append(paras, &cloudformation.Parameter{
		ParameterKey:   aws.String("Subnets"),
		ParameterValue: aws.String("subnet-id-1,subnet-id-2,subnet-id-3"),
	})

	input := &cloudformation.UpdateStackInput{
		Capabilities: aws.StringSlice([]string{"CAPABILITY_IAM"}),
		StackName:    aws.String(stackName),
		Parameters:   paras,
		TemplateBody: aws.String(templateBody),
	}

	result, err := svc.UpdateStack(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

	fmt.Println("Waiting stack update complete ...")

	// snippet-start:[cfn.go.create_stack.wait]
	err = svc.WaitUntilStackUpdateComplete(&cloudformation.DescribeStacksInput{
		StackName: aws.String(stackName),
	})

	if err != nil {
		fmt.Println("Got an error waiting for stack to be created")
		return
	}

	fmt.Println("finish update nodegroup")
}