package queue

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	sfnType "github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/aws-sdk-go/aws"
)

type ListerMock struct {
	errStateMachines       error
	stateMachineExecutions map[string]int
	errExecutions          error
}

func (lm *ListerMock) ListStateMachines(_ context.Context, _ *sfn.ListStateMachinesInput, _ ...func(*sfn.Options)) (*sfn.ListStateMachinesOutput, error) {
	stateMachines := make([]sfnType.StateMachineListItem, len(lm.stateMachineExecutions))
	for i := 0; i < len(lm.stateMachineExecutions); i++ {
		stateMachines[i] = sfnType.StateMachineListItem{StateMachineArn: aws.String(fmt.Sprint(i))}
	}
	return &sfn.ListStateMachinesOutput{
		StateMachines: stateMachines,
	}, lm.errStateMachines
}

func (lm *ListerMock) ListExecutions(_ context.Context, in *sfn.ListExecutionsInput, _ ...func(*sfn.Options)) (*sfn.ListExecutionsOutput, error) {
	executions := make([]sfnType.ExecutionListItem, lm.stateMachineExecutions[*in.StateMachineArn])
	return &sfn.ListExecutionsOutput{
		Executions: executions,
	}, lm.errExecutions
}

func NewListerMock(nStateMachineExecutions []int, errStateMachines, errExecutions error) *ListerMock {

	nex := make(map[string]int, len(nStateMachineExecutions))
	for i, n := range nStateMachineExecutions {
		nex[fmt.Sprint(i)] = n
	}
	return &ListerMock{
		stateMachineExecutions: nex,
		errStateMachines:       errStateMachines,
		errExecutions:          errExecutions,
	}
}

var err1 = fmt.Errorf("error 1")
var err2 = fmt.Errorf("error 2")

func TestGetRunningStateNumber(t *testing.T) {
	cases := []struct {
		input   *ListerMock
		want    int
		wantErr error
	}{
		{
			input:   NewListerMock([]int{}, nil, nil),
			want:    0,
			wantErr: nil,
		},
		{
			input:   NewListerMock([]int{3, 5, 2}, nil, nil),
			want:    10,
			wantErr: nil,
		},
		{
			input:   NewListerMock([]int{}, err1, err2),
			want:    0,
			wantErr: err1,
		},
	}

	for n, cs := range cases {
		got, err := GetRunningStateNumber(cs.input, int32(10))
		if cs.wantErr != err {
			t.Errorf("Case %d: \nGot: %+v\nWant: %+v", n, err, cs.wantErr)
		}
		if got != cs.want {
			t.Errorf("Case %d: \nGot: %d\nWant: %d", n, got, cs.want)
		}
	}

}
