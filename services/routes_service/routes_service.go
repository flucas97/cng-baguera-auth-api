package routes_service

import (
	"io"

	"github.com/flucas97/cng/cng-baguera-auth-api/domain/auth"
	"github.com/flucas97/cng/cng-baguera-auth-api/domain/cannabis"
	"github.com/flucas97/cng/cng-baguera-auth-api/services/cannabis_service"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
)

var (
	RoutesService = &routesService{}
)

type routesServiceInterface interface {
	CallCannabis(method string, url string, body io.ReadCloser, jwt string) (*cannabis.Cannabis, *error_factory.RestErr)
}

type routesService struct{}

// CallCannabis is the entrypoint to Cannabis API
func (rs *routesService) CallCannabis(body io.ReadCloser, jwt string) (*cannabis.Cannabis, *error_factory.RestErr) {
	cannabisRepositoryId, err := auth.GetValueFromJwtKey(jwt, "cannabis_repository_id")
	if err != nil {
		return nil, err
	}

	result, err := cannabis_service.CannabisService.New(body, cannabisRepositoryId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
