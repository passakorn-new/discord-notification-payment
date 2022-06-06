package database

import (
	"context"

	"cloud.google.com/go/firestore"
)

func CreateClient(ctx context.Context) (*firestore.Client, error) {
	projectID := "discord-notification-payment"
	client, err := firestore.NewClient(ctx, projectID)

	return client, err
}
