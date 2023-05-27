package service

import (
	"io/ioutil"
	"net/http"
)

func GetAllUserRepositories(user string) ([]byte, error) {
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