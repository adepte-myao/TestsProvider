package main

import (
	"fmt"
	"os"

	"github.com/adepte-myao/test_provider/internal/handlers"
	"github.com/adepte-myao/test_provider/internal/server"
	"github.com/adepte-myao/test_provider/internal/services"
	"github.com/adepte-myao/test_provider/internal/storage"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load("deploy_task_provider/.env")
	if err != nil {
		fmt.Print(err)
		return
	}

	logger := logrus.New()
	str := os.Getenv("LOG_LEVEL")
	logLevel, err := logrus.ParseLevel(str)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Level = logLevel

	dbConn := os.Getenv("DATABASE_URL")
	db, closeDbFunc, err := storage.NewPostgresDb(dbConn)
	if err != nil {
		logger.Error(err)
		return
	}
	defer closeDbFunc(db)

	taskRepo := storage.NewPostgresTaskRepository(db)
	taskServ := services.NewTaskServiceBase(logger, taskRepo)
	getRandomTopicTaskHandler := handlers.NewGetRandomTopicTaskHandler(logger, taskServ)

	router := mux.NewRouter()
	router.HandleFunc("/get-random-topic-task", getRandomTopicTaskHandler.Handle)

	addr := ":" + os.Getenv("SERVICE_PORT")
	serv := server.NewServer(logger, router, addr)

	err = serv.Start()
	if err != nil {
		logger.Error(err)
		return
	}
}
