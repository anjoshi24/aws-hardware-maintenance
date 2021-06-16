package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	session := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-west-2")}))
	svc := ec2.New(session)

	result, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{
		instance - state - name: []string{"running"},
	})

	if err != nil {
		panic(err.Error())
	}
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("%s", *instance.InstanceId)
		}
	}

	// Create an Amazon EC2 service client
	client := ec2.NewFromConfig(config)

	// DescribeInstanceStatus will list any scheduled maintenance
	result2, err := client.DescribeInstanceStatus(context.TODO(), &ec2.DescribeInstanceStatusInput{
		InstanceIds: []string{"i-03e004d36baecd264"}, //test instance-id passed, list of innstance ids to be passed here.
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Instance Status is :- ")
	for _, x := range result2.InstanceStatuses {
		fmt.Println(*&x.Events)
	}
}
