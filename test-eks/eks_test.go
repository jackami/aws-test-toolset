package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"testing"
)

func Test_CreateNodeGroup(t *testing.T) {
	region := "us-east-1"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := eks.New(sess)
	input := &eks.CreateNodegroupInput{
		//AmiType:            aws.String("AL2_x86_64"),
		ClientRequestToken: nil,
		ClusterName:        aws.String("test-private"),
		//DiskSize:           aws.Int64(20),
		//InstanceTypes:      aws.StringSlice([]string{"t3.medium"}),
		Labels:             nil,
		LaunchTemplate:     &eks.LaunchTemplateSpecification{
			Id:      aws.String("lt-id"),
			//Name:    aws.String("test-eks-template"),
			Version: aws.String("4"),
		},
		NodeRole:           aws.String("arn:aws:iam::AccountID:role/test-private-ng-1"),
		NodegroupName:      aws.String("test-private-ng-4"),
		//ReleaseVersion:     aws.String("1.16.8-20200507"),
		//RemoteAccess:       &eks.RemoteAccessConfig{
		//	Ec2SshKey: aws.String("my_test_ec2_key_pair"),
		//},
		ScalingConfig:      &eks.NodegroupScalingConfig{
			DesiredSize: aws.Int64(2),
			MaxSize:     aws.Int64(2),
			MinSize:     aws.Int64(2),
		},
		Subnets:            aws.StringSlice([]string{"subnet-id-1","subnet-id-2"}),
		Tags:               nil,
		//Version:            aws.String("1.16"),
	}

	result, err := svc.CreateNodegroup(input)
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

func Test_DescribeNodeGroup(t *testing.T) {
	region := "us-east-1"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := eks.New(sess)
	input := &eks.DescribeNodegroupInput{
		ClusterName: aws.String("test-private"),
		NodegroupName: aws.String("test-private-ng-4"),
	}

	result, err := svc.DescribeNodegroup(input)
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

func Test_DeleteNodeGroup(t *testing.T) {
	region := "us-east-1"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := eks.New(sess)
	input := &eks.DeleteNodegroupInput{
		ClusterName:   aws.String("test-private"),
		NodegroupName: aws.String("test-private-ng-4"),
	}

	result, err := svc.DeleteNodegroup(input)
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

