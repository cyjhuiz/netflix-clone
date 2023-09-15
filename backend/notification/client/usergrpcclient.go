package client

import (
	"context"
	"github.com/cyjhuiz/netflix-clone/backend/notification/model"
	"github.com/cyjhuiz/netflix-clone/backend/notification/proto"
	"google.golang.org/grpc"
)

type UserAPIGRPCClient struct {
	UserGRPCClient proto.UserGRPCClient
}

func NewUserAPIGRPCClient() (*UserAPIGRPCClient, error) {
	userAPIURL := "localhost:4001"
	conn, err := grpc.Dial(userAPIURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	userGRPCClient := proto.NewUserGRPCClient(conn)
	return &UserAPIGRPCClient{
		UserGRPCClient: userGRPCClient,
	}, nil
}

func (userAPIGRPCClient *UserAPIGRPCClient) GetUsersByUserIDs(userIDs []int64) ([]*model.User, error) {
	ctx := context.Background()

	getUsersByUserIDsRequest := &proto.GetUsersByUserIDsRequest{
		UserIDs: userIDs,
	}

	response, err := userAPIGRPCClient.UserGRPCClient.GetUsersByUserIDs(ctx, getUsersByUserIDsRequest)
	if err != nil {
		return nil, err
	}

	var convertedUsers []*model.User
	for _, user := range response.Users {
		convertedUser := &model.User{
			UserID:         user.UserID,
			Email:          user.Email,
			SubscriptionID: user.SubscriptionID,
		}
		convertedUsers = append(convertedUsers, convertedUser)
	}

	return convertedUsers, nil
}
