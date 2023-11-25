package app

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Davut97/go-user/repo"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	// Setup
	userJson := `{"email": "fo@bo.com", "firstName": "Foo", "lastName": "Bar", "password": "1234567898"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(userJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	db := repo.NewMockUserRepository(ctrl)
	db.EXPECT().Create(gomock.Any()).Return(repo.User{}, nil)
	app := NewApp(e, db, nil)
	err := app.CreateUser(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, rec.Code)

}

func TestCreateUser500(t *testing.T) {
	ctrl := gomock.NewController(t)
	// Setup
	userJson := `{"email": "fo@bo.com", "firstName": "Foo", "lastName": "Bar", "password": "1234567898"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(userJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	db := repo.NewMockUserRepository(ctrl)
	db.EXPECT().Create(gomock.Any()).Return(repo.User{}, errors.New("error"))
	app := NewApp(e, db, nil)
	err := app.CreateUser(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusInternalServerError, rec.Code)

}

func TestCreateUser400(t *testing.T) {
	ctrl := gomock.NewController(t)
	// Setup
	userJson := `{"email": "fo@bo.com", "firstName": "Foo", "lastName": "Bar", "password": "4"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(userJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	db := repo.NewMockUserRepository(ctrl)

	app := NewApp(e, db, nil)
	err := app.CreateUser(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)

}
