package interfaces

import (
	"main/pkg/domain"
)

type WalletRepository interface {
	CreditToUserWallet(amount float64, walletID int) error
	FindUserIdFromOrderID(id int) (int, error)
	FindWalletIdFromUserID(userId int) (int, error)
	CreateNewWallet(userID int) (int, error)
	GetBalance(WalletID int) (int, error)
	GetHistory(walletID, page, limit int) ([]domain.WalletHistory, error)
}
