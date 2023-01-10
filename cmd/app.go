package main

import (
	"context"
	"fmt"
	Back "github.com/globalskye/RustServerInfo-back-end.git"
	"github.com/globalskye/RustServerInfo-back-end.git/logs"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/handler"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os/signal"

	"net/http"
	"os"
	"time"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()
	if err := start(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}

}
func start(ctx context.Context) error {
	logs.InitLogrus()

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error env variables : %s", err.Error())
	}

	db, err := repository.NewMongoConnect()
	repos := repository.NewRepository(db) // working with db
	services := service.NewService(repos) // business logic
	handlers := handler.NewHandler(services)

	srv := new(Back.Server)
	go func() {
		if err := srv.Run(os.Getenv("APP_HOST"), os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error to running http server : %s", err)
		}
	}()
	fmt.Println(fmt.Sprintf("SERVER WORKING ON http://%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))

	log.Printf("server started")

	<-ctx.Done()

	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
	return err
}
