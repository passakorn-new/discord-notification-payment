package models

type User struct {
	Name             string `firestore:"Name" validate:"required"`
	PhoneNumber      string `firestore:"PhoneNumber" validate:"len=10,regexp=^[0-9]{10}$"`
	PaymentImageLink string `firestore:"PaymentImageLink"`
	DiscordId        string `firestore:"DiscordId"`
}
