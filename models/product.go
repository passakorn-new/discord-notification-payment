package models

type Product struct {
	Name       string  `firestore:"Name" validate:"required"`
	TotalPrice float64 `firestore:"TotalPrice"`
}
