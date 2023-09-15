package api

import (
	"fmt"
	grpc2 "github.com/cyjhuiz/netflix-clone/backend/show/handler/grpc"
	"github.com/cyjhuiz/netflix-clone/backend/show/proto"
	"github.com/cyjhuiz/netflix-clone/backend/show/service"
	"google.golang.org/grpc"
	"net"
)

func RunGRPCAPIServer(listenAddr string, episodeService *service.EpisodeService, favouriteService *service.FavouriteService) error {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)

	episodeGRPCServer := grpc2.NewEpisodeGRPCServer(episodeService)
	proto.RegisterEpisodeGRPCServer(server, episodeGRPCServer)

	favouriteGRPCServer := grpc2.NewFavouriteGRPCServer(favouriteService)
	proto.RegisterFavouriteGRPCServer(server, favouriteGRPCServer)

	fmt.Printf("GRPC API server running on port%s", listenAddr)
	return server.Serve(listener)
}
