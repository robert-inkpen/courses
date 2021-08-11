package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	sqsTypes "github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// Set Up & General Functions
func GetAWSConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}
	return cfg
}

func Environ(key string) string {
	env, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Sprintf("Environment variable, %s, not set.", key))
	}
	return env
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	GEN_ID_KEY  = "GeneratorID"
	ECS_ARN_KEY = "ECSArn"
)

var (
	resources           *dynamodb.Client = dynamodb.NewFromConfig(GetAWSConfig())
	RESOURCE_TABLE_NAME string           = Environ("RESOURCE_TABLE_NAME")
	MissingGeneratorErr error            = fmt.Errorf("can't find generator with that ID")
)

func GetSingleGenerator(ctx context.Context, generatorID GeneratorId) (*types.AttributeValueMemberS, error) {
	// taskDef arn is returned
	result, err := resources.GetItem(ctx,
		&dynamodb.GetItemInput{
			Key: map[string]types.AttributeValue{
				GEN_ID_KEY: asAttr(generatorID),
			},
			TableName: &RESOURCE_TABLE_NAME,
		},
	)
	if err != nil {
		return nil, err
	}
	if len(result.Item) == 0 {
		return nil, MissingGeneratorErr
	}
	final := result.Item[ECS_ARN_KEY].(*types.AttributeValueMemberS)
	return final, nil
}

func asAttr(val interface{}) types.AttributeValue {
	switch v := val.(type) {
	case string:
		return &types.AttributeValueMemberS{Value: v}
	case GeneratorId:
		return &types.AttributeValueMemberS{Value: string(v)}
	default:
		panic(fmt.Sprintf("Can't convert %[1]+v (%[1]T) to attribute", val))
	}
}

func NewTaskFromMessage(msg sqsTypes.Message) (*GeneratorTask, error) {
	gt := GeneratorTask{
		OrigMessage: msg,
	}
	err := json.Unmarshal([]byte(*msg.Body), &gt.Params)
	PanicErr(err)

	return &gt, err
}
