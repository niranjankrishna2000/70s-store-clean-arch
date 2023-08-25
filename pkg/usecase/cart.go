package usecase

import (
	"errors"
	interfaces "main/pkg/repository/interface"
	services "main/pkg/usecase/interface"

	"main/pkg/utils/models"
)

type cartUseCase struct {
	repo                interfaces.CartRepository
	inventoryRepository interfaces.InventoryRepository
	userUseCase         services.UserUseCase
	paymentUseCase      services.PaymentUseCase
}

func NewCartUseCase(repo interfaces.CartRepository, inventoryRepo interfaces.InventoryRepository, userUseCase services.UserUseCase, paymentUseCase services.PaymentUseCase) *cartUseCase {
	return &cartUseCase{
		repo:                repo,
		inventoryRepository: inventoryRepo,
		userUseCase:         userUseCase,
		paymentUseCase:      paymentUseCase,
	}
}

func (i *cartUseCase) AddToCart(user_id, inventory_id int) error {
	//check if the desired product has quantity available
	stock, err := i.inventoryRepository.CheckStock(inventory_id)
	if err != nil {
		return err
	}
	//if available then call userRepository
	if stock <= 0 {
		return errors.New("out of stock")
	}

	//find user cart id
	cart_id, err := i.repo.GetCartId(user_id)
	if err != nil {
		return errors.New("some error in geting user cart")
	}
	//if user has no existing cart create new cart
	if cart_id == 0 {
		cart_id, err = i.repo.CreateNewCart(user_id)
		if err != nil {
			return errors.New("cannot create cart from user")
		}
	}

	//add product to line items
	if err := i.repo.AddLineItems(cart_id, inventory_id); err != nil {
		return errors.New("error in adding products")
	}

	return nil
}

func (i *cartUseCase) CheckOut(userid int) (models.CheckOut, error) {

	address, err := i.repo.GetAddresses(userid)
	if err != nil {
		return models.CheckOut{}, err
	}
	//cartchange
	products, err := i.userUseCase.GetCart(userid, 0, 0)
	if err != nil {
		return models.CheckOut{}, err
	}

	paymentmethods, err := i.paymentUseCase.GetPaymentMethods()
	if err != nil {
		return models.CheckOut{}, err
	}

	var price float64
	for _, v := range products {
		price = price + v.DiscountedPrice
	}

	var checkout models.CheckOut

	checkout.Addresses = address
	checkout.Products = products
	checkout.PaymentMethods = paymentmethods
	checkout.TotalPrice = price

	return checkout, err
}

func (i *cartUseCase) ApplyCoupon(id int, coupon string) (models.CheckOut, error) {
	

	valid,err:=i.repo.ValidateCoupon(coupon)
	if err != nil || !valid{
		return models.CheckOut{}, err
	}

	discountRate,err:=i.repo.GetDiscountRate(coupon)
	if err != nil {
		return models.CheckOut{}, err
	}
	//err:=i.repo.ApplyDiscount(discountRate,id)
	address, err := i.repo.GetAddresses(id)
	if err != nil {
		return models.CheckOut{}, err
	}
	//cartchange
	products, err := i.userUseCase.GetCart(id, 0, 0)
	if err != nil {
		return models.CheckOut{}, err
	}

	paymentmethods, err := i.paymentUseCase.GetPaymentMethods()
	if err != nil {
		return models.CheckOut{}, err
	}

	var price float64
	for _, v := range products {
		price = price + v.DiscountedPrice
	}

	price-=price*(float64(discountRate)/100)
	var checkout models.CheckOut

	checkout.Addresses = address
	checkout.Products = products
	checkout.PaymentMethods = paymentmethods
	checkout.TotalPrice = price

	return checkout, err
}
