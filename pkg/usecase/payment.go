package usecase

import (
	"errors"
	"main/pkg/domain"
	interfaces "main/pkg/repository/interface"
	services "main/pkg/usecase/interface"
)

type paymentUseCase struct {
	paymentRepository interfaces.PaymentRepository
}

func NewPaymentUseCase(paymentRepository interfaces.PaymentRepository) services.PaymentUseCase {
	return &paymentUseCase{
		paymentRepository: paymentRepository,
	}
}

func (p *paymentUseCase) AddNewPaymentMethod(paymentMethod string) error {
	if paymentMethod == " " {
		return errors.New("enter method name")
	}
	if err := p.paymentRepository.AddNewPaymentMethod(paymentMethod); err != nil {
		return err
	}
	return nil
}

func (p *paymentUseCase) RemovePaymentMethod(paymentMethodID int) error {
	if paymentMethodID == 0 {
		return errors.New("enter method id")
	}
	if err := p.paymentRepository.RemovePaymentMethod(paymentMethodID); err != nil {
		return err
	}
	return nil
}

func (p *paymentUseCase) GetPaymentMethods() ([]domain.PaymentMethod,error) {
	paymentMethods,err := p.paymentRepository.GetPaymentMethods()
	if err != nil {
		return []domain.PaymentMethod{},err
	}
	return paymentMethods ,nil
}