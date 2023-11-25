package repo

import (
	"context"
	"fmt"
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetUpDB() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	db, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	dbName := randomdata.Adjective() + "-" + randomdata.Noun()
	fmt.Println(dbName)
	return db.Database(dbName), nil

}

func passwordString(password string) *string {
	return &password
}
func TestCreate(t *testing.T) {
	db, err := SetUpDB()
	require.NoError(t, err)
	collection := db.Collection("users")
	repo := NewMongoUserRepository(collection)
	testUser := User{
		Email:     randomdata.Email(),
		FirstName: randomdata.FirstName(randomdata.RandomGender),
		LastName:  randomdata.LastName(),
		Password:  passwordString(randomdata.StringSample("abcdefghijklmnopqrstuvwxyz")),
	}

	user, err := repo.Create(testUser)

	require.NoError(t, err)
	require.NotEmpty(t, user.ID)
	require.Equal(t, testUser.Email, user.Email)
	require.Equal(t, testUser.FirstName, user.FirstName)
	require.Equal(t, testUser.LastName, user.LastName)
	require.NotEmpty(t, user.Password)

	newUser, error := repo.FindOne(user.ID)
	require.NoError(t, error)
	require.Equal(t, user.ID, newUser.ID)
	require.Equal(t, user.Email, newUser.Email)
	require.Equal(t, user.FirstName, newUser.FirstName)
	require.Equal(t, user.LastName, newUser.LastName)
	require.Equal(t, user.Password, newUser.Password)

	require.True(t, CheckPasswordHash(*testUser.Password, *newUser.Password))

}
