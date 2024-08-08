package data

import (
	"context"
	"errors"
	"time"

	"TaskManager3/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	collection *mongo.Collection
	secretKey  string
}

func NewUserService(db *mongo.Database, secretKey string) *UserService {
	return &UserService{
		collection: db.Collection("users"), // Ensure collection name is 'users'
		secretKey:  secretKey,
	}
}

func (us *UserService) CreateUser(user models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)
	if user.Role == "" {
        user.Role = "user" 
    }
	_, err = us.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserService) AuthenticateUser(username, password string) (string, error) {
	var user models.User
	err := us.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
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

	token, err := us.generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) generateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID.Hex(),
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(us.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (us *UserService) GetUserByID(userID string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = us.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	cursor, err := us.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
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

func (us *UserService) IsAdmin(user *models.User) bool {
	return user.Role == "admin"
}
