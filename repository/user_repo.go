package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/passakorn-new/discord-notification-payment/models"
)

type UserRepository interface {
	GetByDocumentRef(c context.Context, userRef *firestore.DocumentRef) models.User
}

type firestoreUserRepository struct {
	client *firestore.Client
}

func NewUserRepository(client *firestore.Client) UserRepository {
	return &firestoreUserRepository{client: client}
}

func (f firestoreUserRepository) GetByDocumentRef(c context.Context, userRef *firestore.DocumentRef) models.User {
	doc, err := userRef.Get(c)
	if err != nil {
		log.Fatalf("User Not Found: %v", err)
	}

	user := &models.User{}
	doc.DataTo(user)
	return *user
}
