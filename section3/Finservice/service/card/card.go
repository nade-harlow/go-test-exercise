package card

import (
	"github.com/nade-harlow/go-test-exercise/section3/Finservice/dto"
	"github.com/nade-harlow/go-test-exercise/section3/Finservice/pkg/flutterwave"
)

type cardService struct {
	flutterwave *flutterwave.Flutterwave
}

type CardService interface {
	CreateCard(req *dto.CreateCardRequest) (*dto.CreateCardResponse, error)
}

func NewCardService() CardService {
	return &cardService{}
}

func (cs *cardService) CreateCard(req *dto.CreateCardRequest) (*dto.CreateCardResponse, error) {
	res, err := cs.CreateCard(req)
	if err != nil {
		return nil, err
	}

	return &dto.CreateCardResponse{
		ID:         res.ID,
		AmountID:   res.AmountID,
		Amount:     res.Amount,
		Currency:   res.Currency,
		CardPan:    res.CardPan,
		MaskedPan:  res.MaskedPan,
		City:       res.City,
		State:      res.State,
		Address1:   res.Address1,
		ZipCode:    res.ZipCode,
		CVV:        res.CVV,
		Expiration: res.Expiration,
		CardType:   res.CardType,
		NameOnCard: res.NameOnCard,
		CreatedAt:  res.CreatedAt,
		IsActive:   res.IsActive,
	}, nil
}
