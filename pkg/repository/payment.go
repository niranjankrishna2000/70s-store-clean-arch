package repository

import (
	"main/pkg/domain"
	interfaces "main/pkg/repository/interface"

	"gorm.io/gorm"
)

type paymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &paymentRepository{
		DB: DB,
	}
}

func (p *paymentRepository)AddNewPaymentMethod(paymentMethod string)error{
	query:=`INSERT INTO payment_methods(payment_method) VALUES(?)`
	if err:=p.DB.Exec(query,paymentMethod);err!=nil{
		return err.Error
	}
	
	return nil
}

func (p *paymentRepository)RemovePaymentMethod(paymentMethodID int)error{
	query:=`DELETE FROM payment_methods WHERE id=?`
	if err:=p.DB.Exec(query,paymentMethodID);err!=nil{
		return err.Error
	}
	
	return nil
}

func (p *paymentRepository)GetPaymentMethods()([]domain.PaymentMethod,error){
	var paymentMethods []domain.PaymentMethod
	if err:=p.DB.Raw(`SELECT * FROM payment_methods`).Scan(&paymentMethods).Error;err!=nil{
		return []domain.PaymentMethod{},err
	}
	
	return paymentMethods,nil

}