package client

import (
	"context"
	"github.com/cyjhuiz/netflix-clone/backend/notification/model"
	"github.com/cyjhuiz/netflix-clone/backend/notification/proto"
	"google.golang.org/grpc"
)

type ShowAPIGRPCClient struct {
	EpisodeGRPCClient   proto.EpisodeGRPCClient
	FavouriteGRPCClient proto.FavouriteGRPCClient
}

func NewShowAPIGRPCClient() (*ShowAPIGRPCClient, error) {
	showAPIURL := "localhost:4002"
	conn, err := grpc.Dial(showAPIURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	episodeGRPCClient := proto.NewEpisodeGRPCClient(conn)
	favouriteGRPCClient := proto.NewFavouriteGRPCClient(conn)
	return &ShowAPIGRPCClient{
		EpisodeGRPCClient:   episodeGRPCClient,
		FavouriteGRPCClient: favouriteGRPCClient,
	}, nil
}

func (showAPIGRPCClient ShowAPIGRPCClient) GetUserFavouritesByShowID(showID int64) ([]*model.Favourite, error) {
	ctx := context.Background()

	userFavouritesRequest := &proto.GetFavouritesByShowIDRequest{
		ShowID: showID,
	}

	response, err := showAPIGRPCClient.FavouriteGRPCClient.
		GetFavouritesByShowID(
			ctx,
			userFavouritesRequest,
		)

	if err != nil {
		return nil, err
	}

	var convertedUserFavourites []*model.Favourite
	for _, userFavourite := range response.Favourites {
		convertedFavourite := &model.Favourite{
			FavouriteID: userFavourite.FavouriteID,
			ShowID:      userFavourite.ShowID,
			UserID:      userFavourite.UserID,
		}
		convertedUserFavourites = append(convertedUserFavourites, convertedFavourite)
	}

	return convertedUserFavourites, nil
}

func (showAPIGRPCClient *ShowAPIGRPCClient) GetEpisodeByShowIDAndNumber(showID int64, number int64) (*model.Episode, error) {
	ctx := context.Background()

	getEpisodeByShowIDAndNumberRequest := &proto.GetEpisodeByShowIDAndNumberRequest{
		ShowID: showID,
		Number: number,
	}

	response, err := showAPIGRPCClient.EpisodeGRPCClient.GetEpisodeByShowIDAndNumber(ctx, getEpisodeByShowIDAndNumberRequest)
	if err != nil {
		return nil, err
	}
	episode := response

	convertedEpisode := &model.Episode{
		EpisodeID:    episode.EpisodeID,
		ShowID:       episode.ShowID,
		Number:       episode.Number,
		Title:        episode.Title,
		Description:  episode.Description,
		VideoURL:     episode.VideoURL,
		ThumbnailURL: episode.ThumbnailURL,
		ReleaseDate:  episode.ReleaseDate,
		Show: &model.Show{
			ShowID:       episode.Show.ShowID,
			Title:        episode.Show.Title,
			Description:  episode.Show.Description,
			Duration:     episode.Show.Duration,
			ShowType:     episode.Show.ShowType,
			CategoryID:   episode.Show.CategoryID,
			ThumbnailURL: episode.Show.ThumbnailURL,
			ReleaseDate:  episode.Show.ReleaseDate,
			UploaderID:   episode.Show.UploaderID,
		},
	}

	return convertedEpisode, nil
}
