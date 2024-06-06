package handler

import (
	"encoding/json"
	"github.com/nade-harlow/go-test-exercise/section3/Finservice/dto"
	"github.com/nade-harlow/go-test-exercise/section3/Finservice/service/card"
	"io/ioutil"
	"net/http"
)

type cardHandler struct {
	cardService card.CardService
}

func NewCardHandler(card card.CardService) *cardHandler {
	return &cardHandler{
		cardService: card,
	}
}

func (ch *cardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Parse the request body into the CreateCardRequest struct
	var req dto.CreateCardRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Call the card service to create a virtual card
	card, err := ch.cardService.CreateCard(&req)
	if err != nil {
		http.Error(w, "Failed to create virtual card", http.StatusInternalServerError)
		return
	}

	// Return the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(card); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
