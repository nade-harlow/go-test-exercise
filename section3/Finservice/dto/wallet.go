package dto

type CreateWalletResponse struct {
	PublicKey string `json:"address"`
	SecretKey string `json:"password"`
}
