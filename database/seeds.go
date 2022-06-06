package database

import (
	"context"
	"log"
	"strconv"
	"sync"

	"cloud.google.com/go/firestore"
	"github.com/passakorn-new/discord-notification-payment/models"
	"google.golang.org/api/iterator"
)

func DataSeeds(ctx context.Context, client *firestore.Client) {
	var wg sync.WaitGroup
	var wg1 sync.WaitGroup

	wg.Add(3)
	go DeleteCollection(&wg, ctx, client, "users")
	go DeleteCollection(&wg, ctx, client, "products")
	go DeleteCollection(&wg, ctx, client, "product_members")
	wg.Wait()

	wg1.Add(2)
	go CreateUser(&wg1, ctx, client)
	go CreateProduct(&wg1, ctx, client)
	go CreateProductMember(&wg1, ctx, client)
	wg1.Wait()
}

func CreateProductMember(wg *sync.WaitGroup, ctx context.Context, client *firestore.Client) {
	productsMembers := []models.ProductMembers{
		{Product: client.Doc("products/2"), Owner: client.Doc("users/0982758088"), Members: []*firestore.DocumentRef{
			client.Doc("users/0900000000"), client.Doc("users/0982758011"), client.Doc("users/0982799999"), client.Doc("users/0982709008"), client.Doc("users/0982709022"),
		}},
		{Product: client.Doc("products/1"), Owner: client.Doc("users/0982799999"), Members: []*firestore.DocumentRef{
			client.Doc("users/0900000000"), client.Doc("users/0982758011"),
		}},
	}

	for _, productsMember := range productsMembers {
		ref := client.Collection("product_members").NewDoc()
		ref.Set(ctx, productsMember)
	}
}

func CreateUser(wg *sync.WaitGroup, ctx context.Context, client *firestore.Client) {
	users := []models.User{
		{Name: "New", PhoneNumber: "0982758088", DiscordId: "319497873555324928", PaymentImageLink: "https://firebasestorage.googleapis.com/v0/b/discord-notification-payment.appspot.com/o/IMG_1729.JPG?alt=media&token=b6659155-d7aa-432b-8ade-2f1b62364d1e"},
		{Name: "Ken", PhoneNumber: "0900000000", DiscordId: "438346554030555140", PaymentImageLink: ""},
		{Name: "Win", PhoneNumber: "0982758011", DiscordId: "436206867946012682", PaymentImageLink: ""},
		{Name: "Bird", PhoneNumber: "0982799999", DiscordId: "321236200365096961", PaymentImageLink: ""},
		{Name: "Q", PhoneNumber: "0982709008", DiscordId: "764572213742796850", PaymentImageLink: ""},
		{Name: "JaneJi", PhoneNumber: "0982709022", DiscordId: "", PaymentImageLink: ""},
	}

	for _, user := range users {
		client.Collection("users").Doc(user.PhoneNumber).Set(ctx, user)
	}

	defer wg.Done()
}

func CreateProduct(wg *sync.WaitGroup, ctx context.Context, client *firestore.Client) {
	products := []models.Product{
		{Name: "Netflix", TotalPrice: 300},
		{Name: "Youtube Premium", TotalPrice: 299},
	}

	for index, product := range products {
		client.Collection("products").Doc(strconv.Itoa(index+1)).Set(ctx, product)
	}

	defer wg.Done()
}

func DeleteCollection(wg *sync.WaitGroup, ctx context.Context, client *firestore.Client, collectionName string) {
	iter := client.Collection(collectionName).Documents(ctx)
	batch := client.Batch()

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Print(err)
		}

		batch.Delete(doc.Ref)
	}

	_, err := batch.Commit(ctx)
	if err != nil {
		log.Print(err)
	}

	defer iter.Stop()
	defer wg.Done()
}
