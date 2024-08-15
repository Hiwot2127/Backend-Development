package Domain

import(
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskInitialization(t *testing.T){
	id:=primitive.NewObjectID()
	userID:=primitive.NewObjectID()
	now:=time.Now()

	task:= Task{
		ID:          id,
		Title:       "Sample Task",
		Description: "This is a sample task.",
		DueDate:     now,
		Status:      "Pending",
		UserID:      userID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	assert.Equal(t, id, task.ID)
	assert.Equal(t, "Sample Task", task.Title)
	assert.Equal(t, "This is a sample task.", task.Description)
	assert.Equal(t, now, task.DueDate)
	assert.Equal(t, "Pending", task.Status)
	assert.Equal(t, userID, task.UserID)
	assert.Equal(t, now, task.CreatedAt)
	assert.Equal(t, now, task.UpdatedAt)
}

func TestUserrInitialization (t *testing.T){
	id:=primitive.NewObjectID()

	user:= User{
		ID:       id,
		Username: "john_doe",
		Password: "securepassword",
		Role:     "admin",
	}
	assert.Equal(t,id,user.ID)
	assert.Equal(t, "john_doe", user.Username)
	assert.Equal(t, "securepassword", user.Password)
	assert.Equal(t, "admin", user.Role)
}