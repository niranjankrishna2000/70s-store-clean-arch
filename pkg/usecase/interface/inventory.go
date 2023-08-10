package interfaces

import (
	"main/pkg/utils/models"
	
)

type InventoryUseCase interface {
	AddInventory(inventory models.Inventory, image string) (models.InventoryResponse, error)
	UpdateInventory(ProductID int, Stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error

	ShowIndividualProducts(sku string) (models.Inventory, error)
	ListProducts(page int,limit int) ([]models.Inventory, error)
	SearchProducts(key string,page,limit int) ([]models.Inventory, error)
	GetCategoryProducts(catID int,page,limit int) ([]models.Inventory, error)
}
