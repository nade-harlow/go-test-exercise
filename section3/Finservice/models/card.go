package models

type Card struct {
	ID         string `json:"id"`
	AmountID   int    `json:"amount_id"`
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	CardPan    string `json:"card_pan"`
	MaskedPan  string `json:"masked_pan"`
	City       string `json:"city"`
	State      string `json:"state"`
	Address1   string `json:"address_1"`
	ZipCode    string `json:"zip_code"`
	CVV        string `json:"cvv"`
	Expiration string `json:"expiration"`
	CardType   string `json:"card_type"`
	NameOnCard string `json:"name_on_card"`
	CreatedAt  string `json:"created_at"`
	IsActive   bool   `json:"is_active"`
}
