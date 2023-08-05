package usecase

import (
	"fmt"
	domain "main/pkg/domain"
	"main/pkg/helper"
	internal "main/pkg/helper/pdf"
	interfaces "main/pkg/repository/interface"
	services "main/pkg/usecase/interface"
	"main/pkg/utils/models"
	"time"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	userUseCase     services.UserUseCase
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase services.UserUseCase) *orderUseCase {
	return &orderUseCase{
		orderRepository: repo,
		userUseCase:     userUseCase,
	}
}

func (i *orderUseCase) GetOrders(id int) ([]domain.Order, error) {

	orders, err := i.orderRepository.GetOrders(id)
	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil

}

func (i *orderUseCase) OrderItemsFromCart(userid int, order models.Order) (string, error) {

	cart, err := i.userUseCase.GetCart(userid)
	if err != nil {
		return "", err
	}

	var total float64
	for _, v := range cart {
		total = total + v.Total
	}
	var invoiceItems []*internal.InvoiceData
	for _, v := range cart {
		inventory, err := internal.NewInvoiceData(v.ProductName, int(v.Quantity), v.Total)
		if err != nil {
			panic(err)
		}
		invoiceItems = append(invoiceItems, inventory)
	}
	// Create single invoice
	invoice := internal.CreateInvoice("70's Store", "www.70sstore.store", invoiceItems)
	internal.GenerateInvoicePdf(*invoice)
	fmt.Printf("The Total Invoice Amount is: %f", invoice.CalculateInvoiceTotalAmount())

	//COD
	if order.PaymentID == 1 {
		order_id, err := i.orderRepository.OrderItems(userid, order, total)
		if err != nil {
			return "", err
		}

		if err := i.orderRepository.AddOrderProducts(order_id, cart); err != nil {
			return "", err
		}
		cartID, _ := i.userUseCase.GetCartID(userid)
		if err := i.userUseCase.ClearCart(cartID); err != nil {
			return "", err
		}

	} else if order.PaymentID == 2 {
		// razorpay
		order_id, err := i.orderRepository.OrderItems(userid, order, total)
		if err != nil {
			return "", err
		}

		if err := i.orderRepository.AddOrderProducts(order_id, cart); err != nil {
			return "", err
		}
		link := fmt.Sprintf("http://localhost:1243/users/payment/razorpay?id=%d", order_id)
		return link, err
	}

	//wallet

	return "", nil

}

func (i *orderUseCase) CancelOrder(id int) error {

	err := i.orderRepository.CancelOrder(id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) EditOrderStatus(status string, id int) error {

	err := i.orderRepository.EditOrderStatus(status, id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) AdminOrders() (domain.AdminOrdersResponse, error) {

	var response domain.AdminOrdersResponse

	pending, err := i.orderRepository.AdminOrders("PENDING")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	shipped, err := i.orderRepository.AdminOrders("SHIPPED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	delivered, err := i.orderRepository.AdminOrders("DELIVERED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	canceled, err := i.orderRepository.AdminOrders("CANCELED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	response.Canceled = canceled
	response.Pending = pending
	response.Shipped = shipped
	response.Delivered = delivered
	return response, nil

}

func (i *orderUseCase) DailyOrders() (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, time.UTC)
	SalesReport.Orders, _ = i.orderRepository.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0
	for _, v := range SalesReport.Orders {
		total += v.Price
	}
	SalesReport.TotalRevenue = total

	products, err := i.orderRepository.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIDs := helper.FindMostBoughtProduct(products)
	var bestSellers []string
	for _, v := range bestSellerIDs {
		product, err := i.orderRepository.GetProductNameFromID(v)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (i *orderUseCase) WeeklyOrders() (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := endDate.Add(-time.Duration(endDate.Weekday()) * 24 * time.Hour)
	SalesReport.Orders, _ = i.orderRepository.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0
	for _, v := range SalesReport.Orders {
		total += v.Price
	}
	SalesReport.TotalRevenue = total

	products, err := i.orderRepository.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIDs := helper.FindMostBoughtProduct(products)
	var bestSellers []string
	for _, v := range bestSellerIDs {
		product, err := i.orderRepository.GetProductNameFromID(v)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (i *orderUseCase) MonthlyOrders() (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	SalesReport.Orders, _ = i.orderRepository.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0
	for _, v := range SalesReport.Orders {
		total += v.Price
	}
	SalesReport.TotalRevenue = total

	products, err := i.orderRepository.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIDs := helper.FindMostBoughtProduct(products)
	var bestSellers []string
	for _, v := range bestSellerIDs {
		product, err := i.orderRepository.GetProductNameFromID(v)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers
	return SalesReport, nil
}

func (i *orderUseCase) AnnualOrders() (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := time.Date(endDate.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	SalesReport.Orders, _ = i.orderRepository.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0
	for _, v := range SalesReport.Orders {
		total += v.Price
	}
	SalesReport.TotalRevenue = total

	products, err := i.orderRepository.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIDs := helper.FindMostBoughtProduct(products)
	var bestSellers []string
	for _, v := range bestSellerIDs {
		product, err := i.orderRepository.GetProductNameFromID(v)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (i *orderUseCase) CustomDateOrders(dates models.CustomDates) (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := dates.EndDate
	startDate := dates.StartingDate
	SalesReport.Orders, _ = i.orderRepository.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0
	for _, v := range SalesReport.Orders {
		total += v.Price
	}
	SalesReport.TotalRevenue = total

	products, err := i.orderRepository.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIDs := helper.FindMostBoughtProduct(products)
	var bestSellers []string
	for _, v := range bestSellerIDs {
		product, err := i.orderRepository.GetProductNameFromID(v)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}
