package sub

import (
	"context"
	"errors"

	"github.com/workflow-interoperability/activity/grpc"
	"github.com/workflow-interoperability/activity/lib"
)

type SubActivity struct{}

func (SubActivity) GetProperties(ctx context.Context, in *grpc.ActivityGetPropertiesRq) (*grpc.ActivityGetPropertiesRs, error) {
	return &grpc.ActivityGetPropertiesRs{}, nil
}

func (a SubActivity) SetProperties(ctx context.Context, in *grpc.ActivitySetPropertiesRs) (*grpc.ActivityGetPropertiesRs, error) {
	return &grpc.ActivityGetPropertiesRs{}, nil
}
func (SubActivity) CompleteActivity(ctx context.Context, in *grpc.ActivityCompleteActivityRq) (*grpc.ActivityCompleteActivityRs, error) {
	zbcClient, err := lib.ConnectZeebe()
	if err != nil {
		return &grpc.ActivityCompleteActivityRs{}, err
	}
	client := *zbcClient
	switch in.Key {
	case "sync1":
		_, err = client.NewPublishMessageCommand().MessageName("sync2").CorrelationKey("sync2").Send()
		if err != nil {
			return &grpc.ActivityCompleteActivityRs{}, err
		}
	}
	return &grpc.ActivityCompleteActivityRs{}, errors.New("activity with key can not be found")
}
