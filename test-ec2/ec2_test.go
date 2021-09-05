package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
	"testing"
)

func Test_EC2Describe(t *testing.T)  {
	region := "us-east-1"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := ec2.New(sess)
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String("i-instanceid"),
		},
	}

	result, err := svc.DescribeInstances(input)
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

func Test_CreateLaunchTemplate(t *testing.T) {
	region := "us-east-1"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := ec2.New(sess)
	
	input := &ec2.CreateLaunchTemplateInput{
		LaunchTemplateData: &ec2.RequestLaunchTemplateData{
			ImageId:      aws.String("ami-id"),
			InstanceType: aws.String("t2.small"),
			NetworkInterfaces: []*ec2.LaunchTemplateInstanceNetworkInterfaceSpecificationRequest{
				{
					AssociatePublicIpAddress: aws.Bool(true),
					DeviceIndex:              aws.Int64(0),
					Ipv6AddressCount:         aws.Int64(1),
					SubnetId:                 aws.String("subnet-id"),
				},
			},
			TagSpecifications: []*ec2.LaunchTemplateTagSpecificationRequest{
				{
					ResourceType: aws.String("instance"),
					Tags: []*ec2.Tag{
						{
							Key:   aws.String("Name"),
							Value: aws.String("webserver"),
						},
					},
				},
			},
		},
		LaunchTemplateName: aws.String("my-template"),
		VersionDescription: aws.String("WebVersion1"),
	}

	result, err := svc.CreateLaunchTemplate(input)
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

func Test_DescribeLaunchTemplate(t *testing.T) {
	region := "us-west-2"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := ec2.New(sess)
	input := &ec2.DescribeLaunchTemplatesInput{
		LaunchTemplateIds: []*string{
			aws.String("lt-id"),
		},
	}

	result, err := svc.DescribeLaunchTemplates(input)
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

func TestDescribeAutoScalingGroup(t *testing.T) {
	region := "us-west-2"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := autoscaling.New(sess)
	input := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: aws.StringSlice([]string{"eksctl-eks-spot-eksctl-nodegroup"}),
		MaxRecords:            nil,
		NextToken:             nil,
	}

	result, err := svc.DescribeAutoScalingGroups(input)
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

func Test_CreateAutoScalingGroup(t *testing.T) {
	region := "us-west-2"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := autoscaling.New(sess)
	
	mixInstanPolicy := &autoscaling.MixedInstancesPolicy{
		InstancesDistribution: nil,
		LaunchTemplate:        &autoscaling.LaunchTemplate{
			LaunchTemplateSpecification: &autoscaling.LaunchTemplateSpecification{
				LaunchTemplateId:   nil,
				LaunchTemplateName: nil,
				Version:            nil,
			},
			Overrides:                   []*autoscaling.LaunchTemplateOverrides{
				&autoscaling.LaunchTemplateOverrides{
					InstanceType:     nil,
					WeightedCapacity: nil,
				},
				&autoscaling.LaunchTemplateOverrides{
					InstanceType:     nil,
					WeightedCapacity: nil,
				},
			},
		},
	}
	
	input := &autoscaling.CreateAutoScalingGroupInput{
		AutoScalingGroupName:             nil,
		AvailabilityZones:                nil,
		DefaultCooldown:                  nil,
		DesiredCapacity:                  nil,
		HealthCheckGracePeriod:           nil,
		HealthCheckType:                  nil,
		InstanceId:                       nil,
		LaunchConfigurationName:          nil,
		LaunchTemplate:                   nil,
		LifecycleHookSpecificationList:   nil,
		LoadBalancerNames:                nil,
		MaxInstanceLifetime:              nil,
		MaxSize:                          nil,
		MinSize:                          nil,
		MixedInstancesPolicy:             mixInstanPolicy,
		NewInstancesProtectedFromScaleIn: nil,
		PlacementGroup:                   nil,
		ServiceLinkedRoleARN:             nil,
		Tags:                             nil,
		TargetGroupARNs:                  nil,
		TerminationPolicies:              nil,
		VPCZoneIdentifier:                nil,
	}

	result, err := svc.CreateAutoScalingGroup(input)
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

func Test_CreateAutoScalingGroup_Demo(t *testing.T) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	atsvc := autoscaling.New(sess)
	groupName := "test-eks-self-ng"

	overrides := []*autoscaling.LaunchTemplateOverrides{}
	tags := []*autoscaling.Tag{}
	//overrides = append(overrides, &autoscaling.LaunchTemplateOverrides{
	//	InstanceType:     aws.String("m5.xlarge"),
	//	WeightedCapacity: aws.String("1"),
	//})
	overrides = append(overrides, &autoscaling.LaunchTemplateOverrides{
		InstanceType:     aws.String("m5d.xlarge"),
		WeightedCapacity: aws.String("1"),
	})
	overrides = append(overrides, &autoscaling.LaunchTemplateOverrides{
		InstanceType:     aws.String("m4.xlarge"),
		WeightedCapacity: aws.String("1"),
	})
	overrides = append(overrides, &autoscaling.LaunchTemplateOverrides{
		InstanceType:     aws.String("t3.xlarge"),
		WeightedCapacity: aws.String("1"),
	})

	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("Name"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("test-eks-self-ng-spot-Node"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("alpha.eksctl.io/cluster-name"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("eks-spot-eksctl"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("alpha.eksctl.io/eksctl-version"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("0.26.0"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("alpha.eksctl.io/nodegroup-name"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("test-eks-self-ng"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("alpha.eksctl.io/nodegroup-type"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("unmanaged"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("eksctl.cluster.k8s.io/v1alpha1/cluster-name"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("eks-spot-eksctl"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("eksctl.io/v1alpha2/nodegroup-name"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("test-eks-self-ng"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("k8s.io/cluster-autoscaler/eks-spot-eksctl"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("owned"),
	})

	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("k8s.io/cluster-autoscaler/enabled"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("true"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("k8s.io/cluster-autoscaler/node-template/label/aws.amazon.com/spot"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("true"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("k8s.io/cluster-autoscaler/node-template/label/intent"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("apps"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("k8s.io/cluster-autoscaler/node-template/label/lifecycle"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("Ec2Spot"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("k8s.io/cluster-autoscaler/node-template/taint/spotInstance"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("true:PreferNoSchedule"),
	})
	tags = append(tags, &autoscaling.Tag{
		Key:               aws.String("kubernetes.io/cluster/eks-spot-eksctl"),
		PropagateAtLaunch: aws.Bool(true),
		ResourceType:      aws.String("auto-scaling-group"),
		Value:             aws.String("owned"),
	})

	input := &autoscaling.CreateAutoScalingGroupInput{
		DesiredCapacity: aws.Int64(1),
		MinSize:         aws.Int64(0),
		MaxSize:         aws.Int64(5),
		AvailabilityZones: []*string{
			aws.String("us-west-2a"),
			aws.String("us-west-2b"),
			//aws.String("us-west-2c"),
			aws.String("us-west-2d"),
		},
		NewInstancesProtectedFromScaleIn: aws.Bool(false),
		ServiceLinkedRoleARN:             aws.String("arn:aws:iam::AccountID:role/aws-service-role/autoscaling.amazonaws.com/AWSServiceRoleForAutoScaling"),
		Tags:                             tags,
		TerminationPolicies:              []*string{aws.String("OldestLaunchTemplate"), aws.String("OldestInstance")},
		VPCZoneIdentifier:                aws.String("subnet-id-1, subnet-id-2, subnet-id-3"),
		AutoScalingGroupName:             aws.String(groupName),
		MixedInstancesPolicy: &autoscaling.MixedInstancesPolicy{
			InstancesDistribution: &autoscaling.InstancesDistribution{
				OnDemandBaseCapacity:                aws.Int64(0),
				OnDemandPercentageAboveBaseCapacity: aws.Int64(0),
				SpotAllocationStrategy:              aws.String("capacity-optimized"),
			},
			LaunchTemplate: &autoscaling.LaunchTemplate{
				LaunchTemplateSpecification: &autoscaling.LaunchTemplateSpecification{
					LaunchTemplateId: aws.String("lt-id"),
					Version:aws.String("1"),
				},
				Overrides: overrides,
			},
		},
	}

	output, err := atsvc.CreateAutoScalingGroup(input)
	fmt.Println("output", output)
	fmt.Println("err:", err)
}