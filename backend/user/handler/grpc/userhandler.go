package grpc

import (
	"context"
	"github.com/cyjhuiz/netflix-clone/backend/user/proto"
	"github.com/cyjhuiz/netflix-clone/backend/user/service"
)

type UserGRPCServer struct {
	UserService *service.UserService
	proto.UnimplementedUserGRPCServer
}

func NewUserGRPCServer(userService *service.UserService) *UserGRPCServer {
	return &UserGRPCServer{
		UserService: userService,
	}
}

func (userGRPCServer *UserGRPCServer) GetUsersByUserIDs(ctx context.Context, request *proto.GetUsersByUserIDsRequest) (*proto.GetUsersByUserIDsResponse, error) {
	users, err := userGRPCServer.UserService.GetUsersByUserIDs(ctx, request.UserIDs)
	if err != nil {
		return nil, err
	}
	var convertedUsers []*proto.UserView
	for _, user := range users {
		convertedUser := &proto.UserView{
			UserID:         user.UserID,
			Email:          user.Email,
			SubscriptionID: user.SubscriptionID,
		}

		if user.Subscription != nil {
			convertedUser.Subscription = &proto.Subscription{
				SubscriptionID: user.Subscription.SubscriptionID,
				Name:           user.Subscription.Name,
				Price:          user.Subscription.Price,
			}
		}
		convertedUsers = append(convertedUsers, convertedUser)
	}

	response := &proto.GetUsersByUserIDsResponse{
		Users: convertedUsers,
	}

	return response, nil
}
