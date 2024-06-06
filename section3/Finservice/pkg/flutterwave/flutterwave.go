package flutterwave

import (
	"github.com/nade-harlow/go-test-exercise/section3/Finservice/dto"
	"net/http"
	"time"
)

type Flutterwave struct {
	secretKey string
	client    *http.Client
}

func NewFlutterwave(secret string) *Flutterwave {
	return &Flutterwave{
		secretKey: secret,
		client: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   30 * time.Second,
		},
	}
}

func (f *Flutterwave) CreateCard() (dto.CreateCardResponse, error) {

	// This is a mock response
	return dto.CreateCardResponse{
		ID:         "123",
		AmountID:   1,
		Amount:     "1000",
		Currency:   "USD",
		CardPan:    "1234567890",
		MaskedPan:  "123456****",
		City:       "Lagos",
		State:      "Lagos",
		Address1:   "123, Main Street",
		ZipCode:    "100001",
		CVV:        "123",
		Expiration: "12/23",
		CardType:   "VISA",
		NameOnCard: "John Doe",
		CreatedAt:  time.Now().String(),
		IsActive:   true,
	}, nil
}
