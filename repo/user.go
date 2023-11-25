package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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
func (r *MongoUserRepository) FindOne(id primitive.ObjectID) (User, error) {
	ctx := context.Background()
	var user User
	err := r.collection.FindOne(ctx, bson.D{
		{"_id", id},
	}).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (r *MongoUserRepository) Update(user User) (User, error) {
	ctx := context.Background()
	if user.Password != nil {
		hashedPassword, err := HashPassword(*user.Password)
		if err != nil {
			return User{}, err
		}
		user.Password = &hashedPassword
	}
	_, err := r.collection.UpdateOne(ctx, bson.D{
		{"_id", user.ID},
	}, bson.D{
		{"$set", bson.D{
			{"email", user.Email},
			{"firstName", user.FirstName},
			{"lastName", user.LastName},
			{"password", user.Password},
		}},
	})
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *MongoUserRepository) Delete(id primitive.ObjectID) error {
	ctx := context.Background()
	_, err := r.collection.DeleteOne(ctx, bson.D{
		{"_id", id},
	})
	if err != nil {
		return err
	}
	return nil
}