package repository

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/passakorn-new/discord-notification-payment/models"
	"google.golang.org/api/iterator"
)

type ProductRepository interface {
	GetRefDocumentByName(c context.Context, name string) *firestore.DocumentRef
	GetByDocumentRef(c context.Context, productRef *firestore.DocumentRef) models.Product
}

type firestoreProductRepository struct {
	client *firestore.Client
}

func NewProductRepository(client *firestore.Client) ProductRepository {
	return &firestoreProductRepository{client: client}
}

func (f firestoreProductRepository) GetRefDocumentByName(c context.Context, name string) *firestore.DocumentRef {
	iter := f.client.Collection("products").Where("Name", "==", name).Documents(c)
	defer iter.Stop()

	doc, err := iter.Next()
	if err == iterator.Done {
		log.Fatalf("Product Ref Not Found: %v", err)
	}

	return doc.Ref
}

func (f firestoreProductRepository) GetByDocumentRef(c context.Context, productRef *firestore.DocumentRef) models.Product {
	doc, err := productRef.Get(c)
	if err != nil {
		log.Fatalf("Product Not Found: %v", err)
	}

	product := &models.Product{}
	doc.DataTo(&product)
	fmt.Print(product)
	return *product
}
