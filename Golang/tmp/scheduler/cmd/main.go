package main

import (
	"context"
	"errors"
	"fmt"

	"engine/generator/load_balancer/scheduler/pkg/queue"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// Connections
var (
	cfg    = queue.GetAWSConfig()
	sqsCli = sqs.NewFromConfig(cfg)
	sfnCli = sfn.NewFromConfig(cfg)
)

// Constants
const (
	MAX_EXECUTIONS = int32(1000) // for sfn max page size
	MAX_RUNNING    = 500
)

// Errors
var (
	ErrNegativeMachines = errors.New("negative state machine space")
)

//Environment Variables
var (
	INPUT_QUEUE_URL   = queue.Environ("INPUT_QUEUE_URL")
	STATE_MACHINE_ARN = queue.Environ("STATE_MACHINE_ARN")
	FAIL_SNS_URL      = queue.Environ("INPUT_QUEUE_URL") // ! change elsewhere
)

// Import Shortening
var (
	PanicErr = queue.PanicErr
)

// Main Functions
func handler(_ context.Context) error {
	// get running state machines
	stateNumb, err := queue.GetStateMachineExecutions(sfnCli, STATE_MACHINE_ARN)
	PanicErr(err)
	// take diff to find available space
	dif := MAX_RUNNING - stateNumb

	if dif == 0 {
		fmt.Println("No space to run new generators")
		return nil
	}
	if dif < 0 {
		return ErrNegativeMachines
	}
	// retrieve # o/msgs to create new items
	messageList, err := queue.GetMessagesFromQueue(sqsCli, dif, INPUT_QUEUE_URL)
	PanicErr(err)
	fmt.Println(messageList)

	// Send messages to have the Json Unmarshalled
	jsonListSuccess, jsonListFail := queue.GetJsonFromList(messageList)
	fmt.Println(jsonListSuccess)
	fmt.Println(jsonListFail)

	// check Db against genid to get gen url
	//send messages to the state machine to be started

	// go routine to start state machine
	failedGenerators, successGenerators, err := queue.StartStateMachineWithPayload(
		context.TODO(),
		jsonListSuccess,
		sfnCli,
		STATE_MACHINE_ARN)
	if err != nil {
		PanicErr(err) //! Dont pannic b/c case where SM is sent bad value
	}

	errDsmQueue := []error{}
	// delete messages from queue
	for _, message := range successGenerators {
		_, err := queue.DeleteSingleSqsmessage(
			sqsCli,
			message,
			INPUT_QUEUE_URL,
		)
		err = append(errDsmQueue, errDSM)
	}

	//Handle failed messages

	// PanicErr(errDm)

	return nil

}
func main() {
	lambda.Start(handler)
}
