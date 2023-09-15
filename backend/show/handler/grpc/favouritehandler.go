package grpc

import (
	"context"
	"github.com/cyjhuiz/netflix-clone/backend/show/proto"
	"github.com/cyjhuiz/netflix-clone/backend/show/service"
)

type FavouriteGRPCServer struct {
	FavouriteService *service.FavouriteService
	proto.UnimplementedFavouriteGRPCServer
}

func NewFavouriteGRPCServer(favouriteService *service.FavouriteService) *FavouriteGRPCServer {
	return &FavouriteGRPCServer{
		FavouriteService: favouriteService,
	}
}

func (favouriteGRPCServer *FavouriteGRPCServer) GetFavouritesByShowID(ctx context.Context, request *proto.GetFavouritesByShowIDRequest) (*proto.GetFavouritesByShowIDResponse, error) {
	favourites, err := favouriteGRPCServer.FavouriteService.GetFavouritesByShowID(ctx, request.ShowID)
	if err != nil {
		return nil, err
	}

	var convertedFavourites []*proto.Favourite
	for _, favourite := range favourites {
		convertedFavourite := &proto.Favourite{
			FavouriteID: favourite.FavouriteID,
			ShowID:      favourite.ShowID,
			UserID:      favourite.UserID,
		}
		convertedFavourites = append(convertedFavourites, convertedFavourite)
	}

	response := &proto.GetFavouritesByShowIDResponse{
		Favourites: convertedFavourites,
	}

	return response, nil
}
