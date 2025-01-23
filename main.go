package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	logger "github.com/charmbracelet/log"
)

func main() {
	prefix := flag.String("prefix", "thisguymartin-pit", "Stack name prefix to match")
	region := flag.String("region", "", "AWS region (optional, defaults to AWS_REGION env var)")
	delete := flag.Bool("delete", false, "Delete matching stacks")

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
			types.StackStatusImportRollbackComplete,
		},
	})

	if err != nil {
		logger.Fatal("failed to list stacks:", "error", err)
		panic(err)
	}

	stacksToDelete := []types.StackSummary{}

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

			stacksToDelete = append(stacksToDelete, stack)
		} else {
			logger.Debug("Skipping based on Prefix filtering", "StackName", *stack.StackName)
		}

	}

	if *delete && len(stacksToDelete) > 0 {
		logger.Warn(fmt.Sprintf("Are you sure you want to delete %d stacks matching prefix '%s'? (y/N): ", len(stacksToDelete), *prefix))
		var confirm string
		fmt.Scanln(&confirm)

		if strings.ToLower(confirm) == "y" {
			for _, stack := range stacksToDelete {
				logger.Info("Deleting stack", "StackName", *stack.StackName)
				err := deleteStack(context.TODO(), client, *stack.StackName)
				if err != nil {
					logger.Error("Failed to delete stack", "StackName", *stack.StackName, "error", err)
					continue
				}
				logger.Info("Successfully deleted stack", "StackName", *stack.StackName)
			}
		} else {
			logger.Info("Stack deletion cancelled")
		}
	}
}

func deleteStack(ctx context.Context, client *cloudformation.Client, stackName string) error {
	// Create a context with a longer timeout (30 minutes)
	ctx, cancel := context.WithTimeout(ctx, 30*time.Minute)
	defer cancel()

	_, err := client.DeleteStack(ctx, &cloudformation.DeleteStackInput{
		StackName: &stackName,
	})
	if err != nil {
		return fmt.Errorf("failed to delete stack %s: %w", stackName, err)
	}

	// Wait for the stack to be deleted with custom waiter options
	waiter := cloudformation.NewStackDeleteCompleteWaiter(client, func(o *cloudformation.StackDeleteCompleteWaiterOptions) {
		o.MinDelay = time.Second * 15
		o.MaxDelay = time.Second * 30
	})

	err = waiter.Wait(ctx, &cloudformation.DescribeStacksInput{
		StackName: &stackName,
	}, 30*time.Minute) // 30 minute timeout

	if err != nil {
		// Check if the stack is already deleted
		_, descErr := client.DescribeStacks(ctx, &cloudformation.DescribeStacksInput{
			StackName: &stackName,
		})
		if descErr != nil {
			// If we can't find the stack, it's successfully deleted
			if strings.Contains(descErr.Error(), "does not exist") {
				return nil
			}
		}
		return fmt.Errorf("failed to wait for stack deletion %s: %w", stackName, err)
	}

	return nil
}
