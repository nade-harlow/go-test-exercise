package models

type Wallet struct {
	PublicKey string `json:"address"`
	SecretKey string `json:"password"`
}
