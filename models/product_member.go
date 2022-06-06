package models

import "cloud.google.com/go/firestore"

type ProductMembers struct {
	Product *firestore.DocumentRef   `firestore:"Product" validate:"required"`
	Owner   *firestore.DocumentRef   `firestore:"Owner" validate:"required"`
	Members []*firestore.DocumentRef `firestore:"Members" validate:"required"`
}
