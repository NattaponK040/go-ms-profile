cdpackage middleware

import (
	"context"
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/labstack/echo/v4"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Auth() echo.MiddlewareFunc {
	return auth
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		opt := option.WithCredentialsFile("auth-server.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			logrus.Error(err)
			return err
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			logrus.Error(err)
			return err
		}

		a := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(a, "Bearer ", "", 1)
		logrus.Info(idToken)
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			logrus.Error(err)
			return err
		}

		c.Set("token", token)
		return next(c)
	}
}