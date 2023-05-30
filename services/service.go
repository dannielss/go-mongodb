package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dannielss/go-mongodb/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewService(repoCollection *mongo.Collection) ServiceInterface {
	return &service{repoCollection: repoCollection}
}

type ServiceInterface interface {
	GetAllUserRepositories(user string) ([]byte, error)
	SaveRepositories(repos []model.RepoStruct)
}

type service struct {
	repoCollection *mongo.Collection
}

func (s *service) GetAllUserRepositories(user string) ([]byte, error) {
	response, err := http.Get("https://api.github.com/users/"+user+"/repos")

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func (s * service) SaveRepositories(repos []model.RepoStruct) {
	for _, repo := range repos {
		_, err := s.repoCollection.InsertOne(context.TODO(), repo)

		if err != nil {
			fmt.Print("err", err)
		}
	}
}