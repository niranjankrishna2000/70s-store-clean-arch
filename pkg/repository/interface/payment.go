package interfaces

import "main/pkg/domain"

type PaymentRepository interface{
	AddNewPaymentMethod(paymentMethod string)error
	RemovePaymentMethod(paymentMethodID int)error
	GetPaymentMethods()([]domain.PaymentMethod,error)
}