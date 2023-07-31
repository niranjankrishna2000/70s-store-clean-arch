package interfaces

import "main/pkg/domain"

type PaymentUseCase interface{
	AddNewPaymentMethod(paymentMethod string)error
	RemovePaymentMethod(paymentMethodID int)error
	GetPaymentMethods()([]domain.PaymentMethod,error)

}