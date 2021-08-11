package queue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqsTypes "github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SqsReader interface {
	ReceiveMessage(context.Context, *sqs.ReceiveMessageInput, ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)
}

func GetMessagesFromQueue(reader SqsReader, numbToRun int, queueUrl string) ([]sqsTypes.Message, error) {
	messageList := []sqsTypes.Message{}
	var err error

	for counter := 0; counter < numbToRun; {
		message, err := reader.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
			QueueUrl: &queueUrl,
		})
		fmt.Println(err)
		messageList = append(messageList, message.Messages...)
		counter = len(messageList)
	}
	return messageList, err
}

type SqsWriter interface {
	SendMessage(context.Context, *sqs.SendMessageInput, ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

func WriteMsgToQueue(writer SqsWriter, msg *GeneratorTask, queueUrl string) (*sqs.SendMessageOutput, error) {
	return writer.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: msg.OrigMessage.Body,
	})
}

func GetJsonFromList(msg []sqsTypes.Message) ([]*GeneratorTask, []*GeneratorTask) {
	sqsMsgSuccess := []*GeneratorTask{}
	sqsMsgFail := []*GeneratorTask{}

	for i := range msg {
		tempTask, err := NewTaskFromMessage(msg[i])
		if err != nil {
			// UnMarshall Error Catch and skip adding to succes slice			tempTask.Errors = err
			sqsMsgFail = append(sqsMsgFail, tempTask)
			continue
		}
		sqsMsgSuccess = append(sqsMsgSuccess, tempTask)
	}
	return sqsMsgSuccess, sqsMsgFail
}

type SqsDeleter interface {
	DeleteMessage(context.Context, *sqs.DeleteMessageInput, ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error)
}

// delete single message from sqs
func DeleteSingleSqsmessage(deleter SqsDeleter, msg *GeneratorTask, queueUrl string) (*sqs.DeleteMessageOutput, error) {
	return deleter.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
		QueueUrl:      &queueUrl,
		ReceiptHandle: msg.OrigMessage.ReceiptHandle,
	})
}
