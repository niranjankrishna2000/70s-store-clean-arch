package interfaces

import "main/pkg/utils/models"

type OfferUseCase interface {
	AddNewOffer(model models.CreateOffer) error
	MakeOfferExpire(catID int) error
}
