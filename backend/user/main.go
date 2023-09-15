package main

import (
	"github.com/cyjhuiz/netflix-clone/backend/user/api"
	"github.com/cyjhuiz/netflix-clone/backend/user/dao"
	"github.com/cyjhuiz/netflix-clone/backend/user/model"
	"github.com/cyjhuiz/netflix-clone/backend/user/service"
	"github.com/cyjhuiz/netflix-clone/backend/user/util/authutil"
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

	userDao := dao.NewUserDao(store)
	seedDataIfNotExist(userDao)

	userService := service.NewUserService(userDao)

	go api.RunGRPCAPIServer(":4001", userService)

	restApiServer := api.NewRESTAPIServer(":3001", userService)

	restApiServer.Run()
}

func seedDataIfNotExist(userDao *dao.UserDao) {
	users, err := userDao.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	if len(users) > 0 {
		return
	}

	defaultPassword, err := authutil.EncryptPassword("1234")
	if err != nil {
		log.Fatal(err)
	}

	user1 := model.NewUser("bobtan@test.com", defaultPassword)
	user2 := model.NewUser("johncena@test.com", defaultPassword)
	newUsers := []*model.User{user1, user2}
	for _, user := range newUsers {
		err := userDao.CreateUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}
}
