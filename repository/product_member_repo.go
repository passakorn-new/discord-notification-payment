package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/passakorn-new/discord-notification-payment/models"
	"google.golang.org/api/iterator"
)

type ProductMembersRepository interface {
	GetByProductRef(c context.Context, productRef *firestore.DocumentRef) models.ProductMembers
}

type firestoreProductMembersRepository struct {
	client *firestore.Client
}

func NewProductMembersRepository(client *firestore.Client) ProductMembersRepository {
	return &firestoreProductMembersRepository{client: client}
}

func (f firestoreProductMembersRepository) GetByProductRef(c context.Context, productRef *firestore.DocumentRef) models.ProductMembers {
	iter := f.client.Collection("product_members").Where("Product", "==", productRef).Documents(c)
	defer iter.Stop()

	doc, err := iter.Next()
	if err == iterator.Done {
		log.Fatalf("Product Members Not Found: %v", err)
	}

	productMembers := &models.ProductMembers{}
	doc.DataTo(&productMembers)
	return *productMembers
}
