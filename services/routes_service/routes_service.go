package routes_service

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/flucas97/cng/cng-baguera-auth-api/domain/auth"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
)

var (
	RoutesService = &routesService{}
)

type routesServiceInterface interface {
	CallCannabis(method string, url string, body io.ReadCloser, repositoryId string) string
}

type routesService struct{}

// CallCannabis is the entrypoint to Cannabis API, the result is response.body as json stringfy
func (rs *routesService) CallCannabis(method string, url string, body io.ReadCloser, jwt string) (string, *error_factory.RestErr) {
	cannabisRepositoryId, restErr := auth.GetValueFromJwtKey(jwt, "cannabis_repository_id")
	if restErr != nil {
		return "", restErr
	}
	// prepare request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logger.Error("error creating request", err)
		return "", error_factory.NewInternalServerError(err.Error())
	}

	// set headers
	req.Header.Set("Cannabis_repository_id", cannabisRepositoryId)
	req.Header.Set("Content-Type", "application/json")

	// set client
	client := &http.Client{Timeout: time.Second * 10}

	// make request
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("error while making the request", err)
		return "", error_factory.NewInternalServerError(err.Error())
	}
	defer resp.Body.Close()

	// bufferyze it
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	return response, nil
}
