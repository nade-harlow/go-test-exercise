package handler

import (
	"encoding/json"
	"github.com/nade-harlow/go-test-exercise/section3/Finservice/service/wallet"
	"net/http"
)

type walletHandler struct {
	walletService wallet.WalletService
}

func NewWalletHandler(wallet wallet.WalletService) *walletHandler {
	return &walletHandler{
		walletService: wallet,
	}
}

func (ch *walletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	// Call the card service to create a virtual card
	wallet, err := ch.walletService.CreateWallet()
	if err != nil {
		http.Error(w, "Failed to create virtual card", http.StatusInternalServerError)
		return
	}

	// Return the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(wallet); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
