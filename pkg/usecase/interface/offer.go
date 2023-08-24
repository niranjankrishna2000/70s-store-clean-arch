package interfaces

import (
	"main/pkg/domain"
	"main/pkg/utils/models"
)

type OfferUseCase interface {
	AddNewOffer(model models.CreateOffer) error
	MakeOfferExpire(catID int) error
	GetOffers(page,limit int)([]domain.Offer,error)

}
