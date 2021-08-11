package queue

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	sfnTypes "github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/aws-sdk-go/aws"
)

// SFN
type SfnLister interface {
	ListExecutions(context.Context, *sfn.ListExecutionsInput, ...func(*sfn.Options)) (*sfn.ListExecutionsOutput, error)
}

func GetStateMachineExecutions(lister SfnLister, stateMachineArn string) (int, error) {
	StatusFilter := sfnTypes.ExecutionStatusRunning
	executions, err := lister.ListExecutions(context.TODO(), &sfn.ListExecutionsInput{
		StateMachineArn: aws.String(stateMachineArn),
		StatusFilter:    StatusFilter,
	})
	if err != nil {
		return 500, err
	}
	count := len(executions.Executions)

	return count, nil
}

type SfnStarter interface {
	StartExecution(context.Context, *sfn.StartExecutionInput, ...func(*sfn.Options)) (*sfn.StartExecutionOutput, error)
}

func StartStateMachineWithPayload(ctx context.Context, runSlice []*GeneratorTask, sfnCli SfnStarter, arn string) ([]*GeneratorTask, []*GeneratorTask, error) {
	OutputSuccess := []*GeneratorTask{}
	OutputFail := []*GeneratorTask{}
	for _, job := range runSlice {
		singleGenId := job.Params.Window.GeneratorId
		stepPayload, err := GetSingleGenerator(ctx, singleGenId) // Accept GeneratorId
		if err != nil {
			log.Printf("Error getting Step Function Payload: %v", err)
			// Write Error to SQS error queue
			continue
		}
		log.Printf("Starting Step Function: %v", asAttr(stepPayload))

		// start the step function
		// !
		_, err = sfnCli.StartExecution(ctx, &sfn.StartExecutionInput{
			StateMachineArn: aws.String(arn),
			Input:           aws.String(stepPayload.Value),
			Name:            aws.String(addTimeStamp(string(singleGenId))),
		})
		if err != nil {
			log.Printf("Error Starting Step Function %v", err)
			OutputFail = append(OutputFail, job)
			continue
		}
		OutputSuccess = append(OutputSuccess, job)
	}

	return OutputFail, OutputSuccess, nil
}

func addTimeStamp(inValue string) string {
	// day , month, year format
	return fmt.Sprint(inValue, "-", time.Now().UTC().Unix())
}
