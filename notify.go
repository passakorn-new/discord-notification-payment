package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"text/template"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/DisgoOrg/disgohook"
	"github.com/DisgoOrg/disgohook/api"
	"github.com/joho/godotenv"
	"github.com/passakorn-new/discord-notification-payment/database"
	"github.com/passakorn-new/discord-notification-payment/models"
	"github.com/passakorn-new/discord-notification-payment/repository"
)

var reqBody struct {
	ProductName string `json:"product_name"`
	SetUpData   bool   `json:"setup_data"`
}

type TemplateNotify struct {
	Product        models.Product
	Owner          models.User
	Members        []models.User
	CurrentDate    string
	PricePerPerson float64
}

func NotifyDiscord(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Printf("json.NewDecoder: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	godotenv.Load()
	ctx := context.Background()
	firestoreClient, err := database.CreateClient(ctx)
	defer firestoreClient.Close()

	if err != nil {
		log.Printf("Failed to create firebase client: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if reqBody.SetUpData {
		database.DataSeeds(ctx, firestoreClient)
		return
	}

	webhook, err := disgohook.NewWebhookClientByToken(nil, nil, os.Getenv("WEBHOOK_ID")+"/"+os.Getenv("WEBHOOK_TOKEN"))
	if err != nil {
		log.Printf("Failed webhook discord: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	product, productOwner, members := PrepareData(ctx, firestoreClient)
	temp := template.Must(template.ParseFiles(os.Getenv("NOTIFY_TEMPLATE_PATH")))
	tempVar := TemplateNotify{
		Product:        product,
		Owner:          productOwner,
		Members:        members,
		PricePerPerson: math.Ceil(product.TotalPrice / float64(len(members)+1)),
		CurrentDate:    time.Now().Format("02-01-2006"),
	}

	var tpl bytes.Buffer
	if err := temp.Execute(&tpl, tempVar); err != nil {
		log.Fatalf("Failed notify template: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result := tpl.String()
	if _, err = webhook.SendMessage(api.NewWebhookMessageCreateBuilder().
		SetContent(result).
		Build(),
	); err != nil {
		log.Fatalf("failed to send webhook message: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	return
}

func PrepareData(ctx context.Context, client *firestore.Client) (models.Product, models.User, []models.User) {
	productMembersRepo := repository.NewProductMembersRepository(client)
	productRepo := repository.NewProductRepository(client)
	userRepo := repository.NewUserRepository(client)

	productRef := productRepo.GetRefDocumentByName(ctx, reqBody.ProductName)
	product := productRepo.GetByDocumentRef(ctx, productRef)

	productMembers := productMembersRepo.GetByProductRef(ctx, productRef)
	productOwner := userRepo.GetByDocumentRef(ctx, productMembers.Owner)

	members := []models.User{}

	for _, userRef := range productMembers.Members {
		member := models.User{}
		doc, err := userRef.Get(ctx)

		if err != nil {
			log.Panicf("Member %s Not Found: %v", userRef.ID, err)
			break
		}

		doc.DataTo(&member)
		members = append(members, member)
	}

	return product, productOwner, members
}
