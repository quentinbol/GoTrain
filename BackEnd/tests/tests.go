package tests

import (
    "context"
	"testing"

	"GoTrain/BackEnd/router"

    "github.com/stretchr/testify/assert"
    _ "github.com/mattn/go-sqlite3"
)

func Test(t *testing.T) {
	client, err := router.InitializeDB()

	ctxUser := context.Background()
	name:= "Quentin"
	user, err := router.CreateUser(ctxUser, client, name)
	if err != nil {
		t.Fatalf("failed creating user: %v", err)
	}

	content := "Test message content"
	ctxMessage := context.Background()
	message, err := router.CreateMessage(ctxMessage, client, content, user.ID)
	if err != nil {
		t.Fatalf("failed creating message: %v", err)
	}

	assert.Equal(t, content, message.Content)
	assert.Equal(t, name, user.Name)
}