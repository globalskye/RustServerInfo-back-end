package main

import (
	"fmt"
	Back "github.com/globalskye/RustServerInfo-back-end.git"
	"github.com/globalskye/RustServerInfo-back-end.git/logs"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/handler"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logs.InitLogrus()

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error env variables : %s", err.Error())
	}

	db, err := repository.NewMongoConnect()
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	repos := repository.NewRepository(db) // working with db
	services := service.NewService(repos) // business logic
	handlers := handler.NewHandler(services)

	srv := new(Back.Server)
	fmt.Println(fmt.Sprintf("SERVER WORKING ON http://%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
	if err := srv.Run(os.Getenv("APP_HOST"), os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error to running http server : %s", err)
	}
}
