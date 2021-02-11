package client

import (
	"sugar-level-client/models"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"strings"
)

type BackendClient struct {
	ServerAddress string
	UserAgent string
	RestClient *resty.Client
}

func NewFromUrlAddress(url string) (*BackendClient, error) {
	if len(strings.TrimSpace(url)) == 0 {
		return nil, errors.New("url is a mandatory parameter")
	}

	 var backend = &BackendClient {
	 	ServerAddress: url,
	 	UserAgent: "BackendClientUserAgent",
	 	RestClient: resty.New(),
	 }

	return backend, nil
}

func (backend *BackendClient) GetUserData() (*[]models.User, error) {
	var users []models.User

	response, err := backend.getGetRequest().Get(backend.ServerAddress)

	if err == nil && response.StatusCode() == 200 {
		err = json.Unmarshal(response.Body(), &users)
	}

	return &users, err
}

func (backend *BackendClient) getGetRequest() *resty.Request {
	return 	backend.RestClient.R().
		SetHeader("User-Agent", backend.UserAgent).
		SetHeader("Accept", "application/json")
}