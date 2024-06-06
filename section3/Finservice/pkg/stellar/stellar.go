package stellar

import (
	"fmt"
	"github.com/stellar/go/keypair"
	"net/http"
	"time"
)

type stellar struct {
	client *http.Client
}

type KeyPair struct {
	PublicKey string
	SecretKey string
}

func NewStellar(secret string) *stellar {
	return &stellar{
		client: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   30 * time.Second,
		},
	}
}

func CreateStellarKeyPair() (*KeyPair, error) {
	pair, err := keypair.Random()
	if err != nil {
		return nil, fmt.Errorf("error creating stellar pair: %v", err)
	}

	return &KeyPair{PublicKey: pair.Address(), SecretKey: pair.Seed()}, nil
}

// CreateStellarAccount adds the created wallet address to the stellar network
func CreateStellarAccount(addr string) {}
