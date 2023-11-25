package app

import (
	"net/http"

	"github.com/Davut97/go-user/repo"
	"github.com/labstack/echo/v4"
)

type CreateUser struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
}

type CreateUserResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (a *App) CreateUser(c echo.Context) error {
	user := new(CreateUser)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request body", Error: err.Error()})
	}
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request body", Error: err.Error()})
	}
	newUser := repo.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  &user.Password,
	}

	// TODO: Check if user already exists
	createdUser, err := a.userRepo.Create(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, CreateUserResponse{ID: createdUser.ID})
}

func (a *App) Login(c echo.Context) error {
	loginRequest := new(LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request body", Error: err.Error()})
	}
	if err := c.Validate(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request body", Error: err.Error()})
	}

	user, err := a.userRepo.FindByEmail(loginRequest.Email)
	// TODO: Check if user not found
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found", Error: err.Error()})
	}

	if !repo.CheckPasswordHash(loginRequest.Password, *user.Password) {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Invalid credentials", Error: "Invalid credentials"})
	}
	// TODO: Generate JWT token
	return c.JSON(http.StatusOK, nil)

}
