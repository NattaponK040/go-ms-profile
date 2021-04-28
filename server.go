package main

import (
	"github.com/sirupsen/logrus"
	"go-ms-profile/config"
	"go-ms-profile/context"
	"go-ms-profile/repository"
	"go-ms-profile/router"
	"os"
)

func main() {
	s := context.CreateServer(
		"go-upload-ImgProfile",
		config.Config{
			Application: "application",
			Env:         os.Getenv("ENV"),
			Resource:    "resource",
			ParenPath:   "",
		},
		nil,
	)

	if client, err := repository.CreateMongoClient(&s.Config); err != nil {
		logrus.Error(err)
		return
	} else {
		s.MongoRepository = repository.NewMongoRepository(client, &s.Config)
		r := router.NewRoutes(s.Serve, s.MongoRepository)
		r.InitRoute()

		logrus.Fatal(s.Serve.Start(s.GetPort(s.Config.Server.Port)))
	}
}
