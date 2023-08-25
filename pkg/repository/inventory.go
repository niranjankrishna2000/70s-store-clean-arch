package repository

import (
	"errors"
	interfaces "main/pkg/repository/interface"
	"main/pkg/utils/models"
	"strconv"

	"gorm.io/gorm"
)

type inventoryRepository struct {
	DB *gorm.DB
}

func NewInventoryRepository(DB *gorm.DB) interfaces.InventoryRepository {
	return &inventoryRepository{
		DB: DB,
	}
}

func (i *inventoryRepository) AddInventory(inventory models.Inventory, url string) (models.InventoryResponse, error) {
	var inventoryResponse models.InventoryResponse
	query := `
    INSERT INTO inventories (category_id, product_name,description, stock, price, image)
    VALUES (?, ?, ?, ?, ?,?) RETURNING id,stock
	`
	i.DB.Raw(query, inventory.CategoryID, inventory.ProductName, inventory.Description, inventory.Stock, inventory.Price, url).Scan(&inventoryResponse)

	return inventoryResponse, nil

}

func (i *inventoryRepository) CheckInventory(pid int) (bool, error) {
	var k int
	err := i.DB.Raw("SELECT COUNT(*) FROM inventories WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return false, err
	}

	if k == 0 {
		return false, err
	}

	return true, err
}

func (i *inventoryRepository) UpdateInventory(pid int, invData models.UpdateInventory) (models.Inventory, error) {

	// Check the database connection
	if i.DB == nil {
		return models.Inventory{}, errors.New("database connection is nil")
	}

	if invData.CategoryID != 0 {
		if err := i.DB.Exec("UPDATE inventories SET category_id = ? WHERE id= ?", invData.CategoryID, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}
	if invData.ProductName != "" && invData.ProductName !="string"{
		if err := i.DB.Exec("UPDATE inventories SET product_name = ? WHERE id= ?", invData.ProductName, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}
	if invData.Description != "" && invData.Description !="string"{
		if err := i.DB.Exec("UPDATE inventories SET description = ? WHERE id= ?", invData.Description, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}
	if invData.Stock != 0 {
		if err := i.DB.Exec("UPDATE inventories SET stock =  ? WHERE id= ?", invData.Stock, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}

	if invData.Price != 0 {
		if err := i.DB.Exec("UPDATE inventories SET price =  ? WHERE id= ?", invData.Price, pid).Error; err != nil {
			return models.Inventory{}, err
		}
	}
	// Retrieve the update
	var inventory models.Inventory
	if err := i.DB.Raw("SELECT * FROM inventories WHERE id=?", pid).Scan(&inventory).Error; err != nil {
		return models.Inventory{}, err
	}

	return inventory, nil
}

func (i *inventoryRepository) DeleteInventory(inventoryID string) error {
	id, err := strconv.Atoi(inventoryID)
	if err != nil {
		return errors.New("converting into integer not happened")
	}

	result := i.DB.Exec("DELETE FROM inventories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}

func (i *inventoryRepository) ShowIndividualProducts(id string) (models.Inventory, error) {
	pid, error := strconv.Atoi(id)
	if error != nil {
		return models.Inventory{}, errors.New("convertion not happened")
	}
	var product models.Inventory
	err := i.DB.Raw(`
	SELECT
		*
		FROM
			inventories
		
		WHERE
			inventories.id = ?
			`, pid).Scan(&product).Error

	if err != nil {
		return models.Inventory{}, errors.New("error retrieved record")
	}
	return product, nil

}

func (ad *inventoryRepository) ListProducts(page int, limit int) ([]models.Inventory, error) {
	// pagination purpose -
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var productDetails []models.Inventory

	if err := ad.DB.Raw("select id,category_id,product_name,description,stock,price,image from inventories limit ? offset ?", limit, offset).Scan(&productDetails).Error; err != nil {
		return []models.Inventory{}, err
	}

	return productDetails, nil

}

func (i *inventoryRepository) CheckStock(pid int) (int, error) {
	var k int
	if err := i.DB.Raw("SELECT stock FROM inventories WHERE id=?", pid).Scan(&k).Error; err != nil {
		return 0, err
	}
	return k, nil
}

func (i *inventoryRepository) CheckPrice(pid int) (float64, error) {
	var k float64
	err := i.DB.Raw("SELECT price FROM inventories WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return 0, err
	}

	return k, nil
}

func (ad *inventoryRepository) SearchProducts(key string, page, limit int) ([]models.Inventory, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var productDetails []models.Inventory

	query := `
		SELECT *
		FROM inventories 
		WHERE product_name ILIKE '%' || ? || '%' 
		OR description ILIKE '%' || ? || '%'
		limit ? offset ?
	`

	if err := ad.DB.Raw(query, key, key, limit, offset).Scan(&productDetails).Error; err != nil {
		return []models.Inventory{}, err
	}

	return productDetails, nil
}

func (ad *inventoryRepository) GetCategoryProducts(catID int, page, limit int) ([]models.Inventory, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var productDetails []models.Inventory

	query := `
		SELECT *
		FROM inventories 
		WHERE category_id=?
		limit ? offset ?
	`

	if err := ad.DB.Raw(query, catID, limit, offset).Scan(&productDetails).Error; err != nil {
		return []models.Inventory{}, err
	}

	return productDetails, nil
}
