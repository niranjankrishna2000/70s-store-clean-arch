package interfaces

import (
	"main/pkg/utils/models"
	"mime/multipart"
)

type InventoryUseCase interface {
	AddInventory(inventory models.Inventory, image *multipart.FileHeader) (models.InventoryResponse, error)
	UpdateInventory(invID int, invData models.UpdateInventory) (models.Inventory, error)
	DeleteInventory(id string) error

	ShowIndividualProducts(s string) (models.Inventory, error)
	ListProducts(page int, limit int) ([]models.Inventory, error)
	SearchProducts(key string, page, limit int) ([]models.Inventory, error)
	GetCategoryProducts(catID int, page, limit int) ([]models.Inventory, error)
}
