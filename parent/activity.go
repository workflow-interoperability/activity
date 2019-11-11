package parent

import (
	"context"
	"errors"

	"github.com/workflow-interoperability/activity/grpc"
	"github.com/workflow-interoperability/activity/lib"
)

type ParentActivity struct{}

func (ParentActivity) GetProperties(ctx context.Context, in *grpc.ActivityGetPropertiesRq) (*grpc.ActivityGetPropertiesRs, error) {
	return &grpc.ActivityGetPropertiesRs{}, nil
}

func (a ParentActivity) SetProperties(ctx context.Context, in *grpc.ActivitySetPropertiesRs) (*grpc.ActivityGetPropertiesRs, error) {
	return &grpc.ActivityGetPropertiesRs{}, nil
}
func (ParentActivity) CompleteActivity(ctx context.Context, in *grpc.ActivityCompleteActivityRq) (*grpc.ActivityCompleteActivityRs, error) {
	zbcClient, err := lib.ConnectZeebe()
	if err != nil {
		return &grpc.ActivityCompleteActivityRs{}, err
	}
	client := *zbcClient
	switch in.Key {
	case "sync1":
		_, err = client.NewPublishMessageCommand().MessageName("sync1").CorrelationKey("sync1").Send()
		if err != nil {
			return &grpc.ActivityCompleteActivityRs{}, err
		}
	}
	return &grpc.ActivityCompleteActivityRs{}, errors.New("activity with key can not be found")
}
