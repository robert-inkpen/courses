package queue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	sfnTypes "github.com/aws/aws-sdk-go-v2/service/sfn/types"
)

// SFN
type SfnLister interface {
	ListStateMachines(context.Context, *sfn.ListStateMachinesInput, ...func(*sfn.Options)) (*sfn.ListStateMachinesOutput, error)
	ListExecutions(context.Context, *sfn.ListExecutionsInput, ...func(*sfn.Options)) (*sfn.ListExecutionsOutput, error)
}

func GetRunningStateNumber(lister SfnLister, maxExecutions int32) (int, error) {

	// return numb of running state machines
	stateOutput, err := lister.ListStateMachines(context.TODO(), &sfn.ListStateMachinesInput{
		MaxResults: maxExecutions,
	})
	if err != nil {
		return 0, err
	}

	type result struct {
		v int
		e error
	}
	results := make(chan *result)
	defer close(results)
	for _, sm := range stateOutput.StateMachines {
		go func(sm sfnTypes.StateMachineListItem) {
			excOut, err := lister.ListExecutions(context.TODO(), &sfn.ListExecutionsInput{
				MaxResults:      maxExecutions,
				StateMachineArn: sm.StateMachineArn,
			})
			if err != nil {
				results <- &result{0, err}
				return
			}
			results <- &result{len(excOut.Executions), nil}
		}(sm)
	}
	var counter int
	for i := 0; i < len(stateOutput.StateMachines); i++ {
		res := <-results
		PanicErr(res.e)
		counter += res.v
	}
	return counter, nil
}
