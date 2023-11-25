package app

import (
	"github.com/labstack/echo/v4"
)

var (
	JokeLimit = 10
)

// get 10 jokes
func (a *App) GetJokes(c echo.Context) error {

	jokes, err := a.joke.GetJokes(JokeLimit)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, jokes)
}
