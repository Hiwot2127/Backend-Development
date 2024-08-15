package Repositories

import (
	"testing"

	"TaskManager5/Domain"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserRepository(t *testing.T) {
	// Setup mtest with MongoDB
	mt := mtest.New(t, mtest.NewOptions().DatabaseName("testdb"))
	defer mt.Close()

	// Create UserRepository instance
	secretKey := "supersecretkey"
	repo := NewUserRepository(mt.DB, secretKey)

	// Test CreateUser
	t.Run("CreateUser", func(t *testing.T) {
		user := Domain.User{
			Username: "testuser",
			Password: "password123",
			Role:     "user",
		}

		createdUser, err := repo.CreateUser(user)
		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, user.Username, createdUser.Username)
		assert.NotEmpty(t, createdUser.Password)

		// Check if password is hashed
		err = bcrypt.CompareHashAndPassword([]byte(createdUser.Password), []byte(user.Password))
		assert.NoError(t, err)
	})

	// Test AuthenticateUser
	t.Run("AuthenticateUser", func(t *testing.T) {
		user := Domain.User{
			Username: "authuser",
			Password: "authpassword",
		}

		_, err := repo.CreateUser(user)
		assert.NoError(t, err)

		token, err := repo.AuthenticateUser(user.Username, user.Password)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	// Test GetUserByID
	t.Run("GetUserByID", func(t *testing.T) {
		user := Domain.User{
			Username: "getuser",
			Password: "getpassword",
		}
		createdUser, err := repo.CreateUser(user)
		assert.NoError(t, err)

		fetchedUser, err := repo.GetUserByID(createdUser.ID.Hex())
		assert.NoError(t, err)
		assert.NotNil(t, fetchedUser)
		assert.Equal(t, user.Username, fetchedUser.Username)
	})

	// Test GetAllUsers
	t.Run("GetAllUsers", func(t *testing.T) {
		user1 := Domain.User{
			Username: "user1",
			Password: "password1",
		}
		user2 := Domain.User{
			Username: "user2",
			Password: "password2",
		}
		_, err := repo.CreateUser(user1)
		assert.NoError(t, err)
		_, err = repo.CreateUser(user2)
		assert.NoError(t, err)

		users, err := repo.GetAllUsers()
		assert.NoError(t, err)
		assert.Len(t, users, 2)
	})

	// Test AuthenticateUser with invalid credentials
	t.Run("AuthenticateUserInvalid", func(t *testing.T) {
		_, err := repo.AuthenticateUser("nonexistentuser", "wrongpassword")
		assert.Error(t, err)
		assert.Equal(t, "invalid username or password", err.Error())
	})
}
