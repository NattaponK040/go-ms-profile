package controller

import (
	"github.com/labstack/echo/v4"
	"go-ms-profile/repository"
	"go-ms-profile/service"
)

type Controller struct {
	mongoRepository *repository.MongoRepository
	profileService  *service.Service
}

func NewController(mg *repository.MongoRepository) *Controller {
	return &Controller{
		mongoRepository: mg,
		profileService:  service.NewProfileService(mg),
	}
}

func (c *Controller) GetProfileAccessTokenController(ctx echo.Context) error {
	return c.profileService.GetProfile()
}

func (c *Controller) GetFacebookAccessTokenController(ctx echo.Context) error {
	return c.profileService.GetFacebookAccessToken()
}

func (c *Controller) GetGoogleAccessTokenController(ctx echo.Context) error {
	return c.profileService.GetGoogleAccessToken()
}

func (c *Controller) GetProfileSetting(ctx echo.Context) error {
	return c.profileService.GetProfileSetting()
}