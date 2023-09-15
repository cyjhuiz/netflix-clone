package main

import (
	"github.com/cyjhuiz/netflix-clone/backend/show/api"
	"github.com/cyjhuiz/netflix-clone/backend/show/dao"
	"github.com/cyjhuiz/netflix-clone/backend/show/service"
	"log"
)

func main() {
	store, err := dao.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	showDao := dao.NewShowDao(store)
	episodeDao := dao.NewEpisodeDao(store)
	likeDao := dao.NewLikeDao(store)
	favouriteDao := dao.NewFavouriteDao(store)
	redisDao := dao.NewRedisDao() // cache store
	redisDao.Flush()

	showService := service.NewShowService(showDao, redisDao)
	episodeService := service.NewEpisodeService(episodeDao)
	likeService := service.NewLikeService(likeDao)
	favouriteService := service.NewFavouriteService(favouriteDao)

	go api.RunGRPCAPIServer(":4002", episodeService, favouriteService)

	restApiServer := api.NewRESTAPIServer(":3002", showService, episodeService, likeService, favouriteService)
	restApiServer.Run()
}
