package queue

import (
	sqsTypes "github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type (
	WindowId        string
	GeneratorId     string
	ViewId          string
	AppName         string
	WindowIdAppName string
	ReceiptHandle   *string
	WindowFromQue   map[WindowIdAppName]Window
	SqsMessage      map[string]InputMessage
)
type (
	Window struct {
		WindowId    WindowId    `json:"window_id"`
		GeneratorId GeneratorId `json:"generator_id"`
		ViewId      ViewId      `json:"view_id"`
		AppName     AppName     `json:"app_name,omitempty"`
	}
	TaskParams struct {
		Window  Window  `json:"window_id"`
		AppName AppName `json:"app_name"`
	}
	InputMessage struct {
		Body          *string
		ReceiptHandle ReceiptHandle
		Id            *string
	}
	GeneratorTask struct {
		Params      TaskParams
		OrigMessage sqsTypes.Message
		Errors      error
	}
)
