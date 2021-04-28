package router

import (
	"github.com/labstack/echo/v4"
	"go-ms-profile/controller"
	"go-ms-profile/middleware"
	"go-ms-profile/repository"
)

type Routes struct {
	app                *echo.Echo
	mg                 *repository.MongoRepository
	controller         *controller.Controller
	firebaseMiddleware echo.MiddlewareFunc
}

func NewRoutes(app *echo.Echo, mg *repository.MongoRepository) *Routes {
	return &Routes{
		app:                app,
		mg:                 mg,
		controller:         controller.NewController(mg),
		firebaseMiddleware: middleware.Auth(),
	}
}
func (r *Routes) InitRoute() {
	route := r.app.POST
	c := r.controller
	route("/getProfile", c.GetProfileAccessTokenController, r.firebaseMiddleware)
	route("/getFacebookAccessToken", c.GetFacebookAccessTokenController, r.firebaseMiddleware)
	route("/getGoogleAccessToken", c.GetGoogleAccessTokenController, r.firebaseMiddleware)
	route("/getProfileSettings", c.GetProfileSetting, r.firebaseMiddleware)
}
