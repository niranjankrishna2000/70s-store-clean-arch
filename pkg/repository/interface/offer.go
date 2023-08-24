package interfaces

import (
	"main/pkg/domain"
	"main/pkg/utils/models"
)

type OfferRepository interface {
	AddNewOffer(offer models.CreateOffer) error
	MakeOfferExpire(catID int) error
	FindDiscountPercentage(catID int) (int, error)
	GetOffers(page,limit int) ([]domain.Offer,error)
}
