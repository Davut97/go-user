package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Password  *string            `json:"password,omitempty"`
}
type UserRepository interface {
	FindOne(id string) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(id string) error
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{collection: collection}
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (r *MongoUserRepository) Create(user User) (User, error) {
	ctx := context.Background()
	hashedPassword, err := HashPassword(*user.Password)
	if err != nil {
		return User{}, err
	}
	user.Password = &hashedPassword
	doc, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return User{}, err
	}
	user.ID = doc.InsertedID.(primitive.ObjectID)
	return user, nil
}
