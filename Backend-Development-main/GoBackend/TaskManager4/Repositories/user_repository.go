package Repositories

import (
	"context"
	"errors"

	"TaskManager4/Domain"
	"TaskManager4/Infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(user Domain.User) (*Domain.User, error)
	AuthenticateUser(username, password string) (string, error)
	GetUserByID(userID string) (*Domain.User, error)
	GetAllUsers() ([]Domain.User, error)
}

type userRepository struct {
	collection *mongo.Collection
	secretKey string
}

func NewUserRepository(db *mongo.Database, secretKey string) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
		secretKey: secretKey,
	}
}

func (ur *userRepository) CreateUser(user Domain.User) (*Domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)
	if user.Role == "" {
		user.Role = "user"
	}
	_, err = ur.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) AuthenticateUser(username, password string) (string, error) {
	var user Domain.User
	err := ur.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("invalid username or password")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := Infrastructure.GenerateJWT(user, ur.secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (ur *userRepository) GetUserByID(userID string) (*Domain.User, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	var user Domain.User
	err = ur.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) GetAllUsers() ([]Domain.User, error) {
	cursor, err := ur.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []Domain.User
	for cursor.Next(context.Background()) {
		var user Domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
