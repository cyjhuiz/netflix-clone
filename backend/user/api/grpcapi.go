package api

import (
	"fmt"
	grpc2 "github.com/cyjhuiz/netflix-clone/backend/user/handler/grpc"
	"github.com/cyjhuiz/netflix-clone/backend/user/proto"
	"github.com/cyjhuiz/netflix-clone/backend/user/service"
	"google.golang.org/grpc"
	"net"
)

func RunGRPCAPIServer(listenAddr string, userService *service.UserService) error {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	var options []grpc.ServerOption
	server := grpc.NewServer(options...)

	userGRPCServer := grpc2.NewUserGRPCServer(userService)

	proto.RegisterUserGRPCServer(server, userGRPCServer)

	fmt.Printf("GRPC API server running on port%s", listenAddr)
	return server.Serve(listener)
}
