package wallet

import (
	"fmt"
	"github.com/nade-harlow/go-test-exercise/section3/Finservice/dto"
	"github.com/nade-harlow/go-test-exercise/section3/Finservice/pkg/stellar"
)

type walletService struct {
}

type WalletService interface {
	CreateWallet() (*dto.CreateWalletResponse, error)
}

func NewWalletService() WalletService {
	return &walletService{}
}

func (w *walletService) CreateWallet() (*dto.CreateWalletResponse, error) {
	stellarPair, err := stellar.CreateStellarKeyPair()
	if err != nil {
		return nil, fmt.Errorf("wallet service error: %v", err)
	}

	return &dto.CreateWalletResponse{
		PublicKey: stellarPair.PublicKey,
		SecretKey: stellarPair.SecretKey,
	}, nil
}
