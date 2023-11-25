package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	Email     string  `json:"email"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Password  *string `json:"-"`
}

type UserRepository interface {
	FindByEmail(email string) (User, error)
	FindOne(id string) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(id string) error
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) (*MongoUserRepository, error) {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)

	return &MongoUserRepository{collection: collection}, err
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
	user.ID = doc.InsertedID.(primitive.ObjectID).Hex()
	return user, nil
}

func (r *MongoUserRepository) FindOne(id string) (User, error) {
	ctx := context.Background()
	var user User
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, err
	}
	err = r.collection.FindOne(ctx, bson.D{
		{"_id", objID},
	}).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *MongoUserRepository) FindByEmail(email string) (User, error) {
	ctx := context.Background()
	var user User
	err := r.collection.FindOne(ctx, bson.D{
		{"email", email},
	}).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *MongoUserRepository) Delete(id string) error {
	ctx := context.Background()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.D{
		{"_id", objID},
	})
	if err != nil {
		return err
	}
	return nil
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
	objID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return User{}, err
	}
	_, err = r.collection.UpdateOne(ctx, bson.D{
		{"_id", objID},
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
