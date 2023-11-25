package app

import (
	"net/http"

	"github.com/Davut97/go-user/repo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type App struct {
	e        *echo.Echo
	db       *mongo.Database
	log      *zap.Logger
	userRepo repo.UserRepository
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewApp(e *echo.Echo, userRepo repo.UserRepository, log *zap.Logger) *App {
	e.Validator = &CustomValidator{validator: validator.New()}
	app := &App{e: e, log: log, userRepo: userRepo}
	app.RegisterRoutes()
	return app

}

func (a *App) Start(port string) error {
	a.log.Info("Starting the server...")
	return a.e.Start(port)
}

func (a *App) Stop() error {
	a.log.Info("Shutting down the server...")
	return a.e.Close()
}
