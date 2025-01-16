package main

import (
	"context"
	"flag"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	logger "github.com/charmbracelet/log"
)

func main() {
	prefix := flag.String("prefix", "thisguymartin-pit", "Stack name prefix to match")
	region := flag.String("region", "", "AWS region (optional, defaults to AWS_REGION env var)")

	flag.Parse()

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region))
	if err != nil {
		logger.Fatal("unable to load SDK config", "error", err)
		panic(err)
	}

	client := cloudformation.NewFromConfig(cfg)
	output, err := client.ListStacks(context.TODO(), &cloudformation.ListStacksInput{
		StackStatusFilter: []types.StackStatus{
			types.StackStatusCreateComplete,
			types.StackStatusCreateInProgress,
			types.StackStatusRollbackComplete,
			types.StackStatusRollbackFailed,
			types.StackStatusUpdateRollbackFailed,
			types.StackStatusUpdateComplete,
		},
	})

	if err != nil {
		logger.Fatal("failed to list stacks:", "error", err)
		panic(err)
	}

	for _, stack := range output.StackSummaries {
		if strings.HasPrefix(*stack.StackName, *prefix) {
			logger.Info("Stack Name: ", "StackName", *stack.StackName)
			logger.Info("Stack ID: ", "StackId", *stack.StackId)
			logger.Info("Stack Status: ", "StackStatus", stack.StackStatus)
			logger.Info("Creation Time: ", "CreationTime", stack.CreationTime)
			if stack.LastUpdatedTime != nil {
				logger.Info("Last Updated: ", "LastUpdatedTime", stack.LastUpdatedTime)
			}
			logger.Info("-------------------")
		}
	}

}
