package main

import (
	"encoding/json"
	"log"

	"github.com/dannielss/go-mongodb/config"
	"github.com/dannielss/go-mongodb/model"
	service "github.com/dannielss/go-mongodb/services"
)


func main() {
	repoCollection, err := config.Init()

	if err != nil {
		log.Fatal("err", err)
	}

	service := service.NewService(repoCollection)

	var repos []model.RepoStruct

	res, err := service.GetAllUserRepositories("dannielss")

	if err != nil {
		log.Fatal("err", err)
	}

	err = json.Unmarshal(res, &repos)

	if err != nil {
		log.Fatal("err", err)
	}

	service.SaveRepositories(repos)
}