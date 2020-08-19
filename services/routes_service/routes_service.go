package routes_service

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/flucas97/cng/cng-baguera-auth-api/domain/auth"
	"github.com/flucas97/cng/cng-baguera-auth-api/domain/cannabis"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
)

var (
	RoutesService = &routesService{}
)

type routesServiceInterface interface {
	CallCannabis(method string, url string, body io.ReadCloser, jwt string) (*cannabis.Cannabis, *error_factory.RestErr)
}

type routesService struct{}

// CallCannabis is the entrypoint to Cannabis API, the result is response.body as json stringfy
func (rs *routesService) CallCannabis(method string, url string, body io.ReadCloser, jwt string) (*cannabis.Cannabis, *error_factory.RestErr) {
	cannabisRepositoryId, restErr := auth.GetValueFromJwtKey(jwt, "cannabis_repository_id")
	if restErr != nil {
		return nil, restErr
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logger.Error("error creating request", err)
		return nil, error_factory.NewInternalServerError(err.Error())
	}

	req.Header.Set("Cannabis_repository_id", cannabisRepositoryId)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("error while making the request", err)
		return nil, error_factory.NewInternalServerError(err.Error())
	}
	defer resp.Body.Close()

	result := cannabis.Cannabis{}

	e := json.NewDecoder(resp.Body).Decode(&result)
	if e != nil {
		logger.Error(e.Error(), e)
		return nil, error_factory.NewInternalServerError(e.Error())
	}

	return &result, nil
}
