package usecase

import (
	interfaces "main/pkg/repository/interface"
	services "main/pkg/usecase/interface"
	"main/pkg/utils/models"
)

type offerUseCase struct {
	offerRepo interfaces.OfferRepository
}

func NewOfferUseCase(repo interfaces.OfferRepository) services.OfferUseCase {
	return &offerUseCase{
		offerRepo: repo,
	}
}

func (o *offerUseCase) AddNewOffer(model models.CreateOffer) error {
	if err := o.offerRepo.AddNewOffer(model); err != nil {
		return err
	}

	return nil
}
func (o *offerUseCase) MakeOfferExpire(catID int) error {
	if err := o.offerRepo.MakeOfferExpire(catID); err != nil {
		return err
	}

	return nil
}
