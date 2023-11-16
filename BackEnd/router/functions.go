package router

import (
	"context"
    "log"
	"fmt"
	"time"
    "GoTrain/BackEnd/ent"

)

func InitializeDB() (*ent.Client, error) {
    client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        return nil, err
    }
    if err := client.Schema.Create(context.Background()); err != nil {
        return nil, err
    }
    return client, nil
}

func CreateUser(ctx context.Context, client *ent.Client, name string) (*ent.Users, error) {
    newUser, err := client.Users.
        Create().
        SetName(name).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }
    return newUser, nil
}

func CreateMessage(ctx context.Context, client *ent.Client, content string, userID int) (*ent.Messages, error) {
	user, err := client.Users.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	message, err := client.Messages.
		Create().
		SetContent(content).
		SetCreatedAt(time.Now()).
		SetUserID(userID).
		SetUser(user).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating message: %w", err)
	}

	log.Println("message was created: ", message)
	return message, nil
}

func QueryAllMessages(ctx context.Context, client *ent.Client) ([]*ent.Messages, error) {
    messages, err := client.Messages.
        Query().
        All(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed querying messages: %w", err)
    }
    log.Println("messages returned: ", messages)
    return messages, nil
}
