package main

import (
	"encoding/json"
	"fmt"

	"github.com/dannielss/go-mongodb/model"
	service "github.com/dannielss/go-mongodb/services"
)

func main() {
	var repos []model.RepoStruct

	res, err := service.GetAllUserRepositories("dannielss");

	if err != nil {
		fmt.Print("err", err)
	}

	err = json.Unmarshal(res, &repos)

	if err != nil {
		fmt.Print("err", err)
	}
}