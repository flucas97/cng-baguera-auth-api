package cannabis_service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/flucas97/cng/cng-baguera-auth-api/domain/cannabis"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
)

const (
	cannabisBaseUrl = "http://172.30.0.8:8083/api"
)

var CannabisService cannabisServiceInterface = &cannabisService{}

type cannabisService struct{}

type cannabisServiceInterface interface {
	New(body io.ReadCloser, repositoryId string) (*cannabis.Cannabis, *error_factory.RestErr)
}

func (cs cannabisService) New(body io.ReadCloser, repositoryId string) (*cannabis.Cannabis, *error_factory.RestErr) {
	var (
		cannabis = cannabis.Cannabis{}
	)

	jsonResponse, err := makeRequest(repositoryId, "POST", "/new-cannabis", body)
	if err != nil {
		return nil, err
	}

	e := json.Unmarshal([]byte(jsonResponse), &cannabis)
	if e != nil {
		return nil, error_factory.NewInternalServerError(e.Error())
	}

	return &cannabis, nil
}

func GetCannabis() {

}

func GetAllCannabis(jwt string, repositoryId string) ([]cannabis.Cannabis, *error_factory.RestErr) {
	var (
		cannabis = []cannabis.Cannabis{}
	)

	jsonResponse, err := makeRequest(repositoryId, "GET", "cannabis", nil)
	if err != nil {
		return nil, err
	}

	e := json.Unmarshal([]byte(jsonResponse), &cannabis)
	if e != nil {
		return nil, error_factory.NewInternalServerError(e.Error())
	}
	return cannabis, nil
}

func UpdateCannabis() {

}

func DeleteCannabis() {

}

func makeRequest(repositoryId string, method string, path string, body io.ReadCloser) ([]byte, *error_factory.RestErr) {
	req, err := http.NewRequest(method, cannabisBaseUrl+path, body)
	if err != nil {
		logger.Error("error creating request", err)
		return nil, error_factory.NewInternalServerError(err.Error())
	}

	req.Header.Set("Cannabis-repository-id", repositoryId)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("error while making the request", err)
		return nil, error_factory.NewInternalServerError(err.Error())
	}
	defer resp.Body.Close()

	jsonResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("error while reading the request", err)
		return nil, error_factory.NewInternalServerError(err.Error())
	}
	return jsonResponse, nil
}
