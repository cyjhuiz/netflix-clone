package grpc

import (
	"context"
	"github.com/cyjhuiz/netflix-clone/backend/show/proto"
	"github.com/cyjhuiz/netflix-clone/backend/show/service"
)

type EpisodeGRPCServer struct {
	EpisodeService *service.EpisodeService
	proto.UnimplementedEpisodeGRPCServer
}

func NewEpisodeGRPCServer(episodeService *service.EpisodeService) *EpisodeGRPCServer {
	return &EpisodeGRPCServer{
		EpisodeService: episodeService,
	}
}

func (episodeGRPCServer *EpisodeGRPCServer) GetEpisodeByShowIDAndNumber(ctx context.Context, request *proto.GetEpisodeByShowIDAndNumberRequest) (*proto.GetEpisodeByShowIDAndNumberResponse, error) {
	episode, err := episodeGRPCServer.EpisodeService.GetEpisodeByShowIDAndNumber(ctx, request.ShowID, request.Number)
	if err != nil {
		return nil, err
	}

	response := &proto.GetEpisodeByShowIDAndNumberResponse{
		EpisodeID:    episode.EpisodeID,
		ShowID:       episode.ShowID,
		Number:       episode.Number,
		Title:        episode.Title,
		Description:  episode.Description,
		VideoURL:     episode.VideoURL,
		ThumbnailURL: episode.ThumbnailURL,
		ReleaseDate:  episode.ReleaseDate,
		Show: &proto.Show{
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

	return response, nil
}
