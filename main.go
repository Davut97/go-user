package main

import (
	"context"
	"net/http"
	"time"

	"github.com/Davut97/go-user/pkg/app"
	"github.com/Davut97/go-user/pkg/config"
	"github.com/Davut97/go-user/pkg/joke"
	"github.com/Davut97/go-user/repo"
	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	cn, err := config.GetConfig()
	if err != nil {
		logger.Error("Failed to get config", zap.Error(err))
		return
	}
	//"https://api.chucknorris.io"
	httpClient := http.Client{}
	httpClient.Timeout = time.Second * time.Duration(cn.JokesTimeout)
	jokeClient := joke.NewChuckNorrisJokeClient(cn.JokesURL, &httpClient, cn.JokesLimit)

	//("mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(cn.DBConnectionString)
	db, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logger.Error("Failed to connect to MongoDB", zap.Error(err))
		return
	}
	userRepo, err := repo.NewMongoUserRepository(db.Database(cn.DBName).Collection("users"))
	if err != nil {
		logger.Error("Failed to create user repository", zap.Error(err))
		return
	}
	e := echo.New()
	if err != nil {
		logger.Error("Failed to create echo instance", zap.Error(err))
		return
	}
	a := app.NewApp(e, userRepo, logger, jokeClient)
	if err := a.Start(cn.BindAddress); err != nil {
		logger.Error("Failed to start server", zap.Error(err))
		return
	}

}
