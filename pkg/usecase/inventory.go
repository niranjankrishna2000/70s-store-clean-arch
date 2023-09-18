package usecase

import (
	"errors"
	"fmt"
	"main/pkg/helper"
	interfaces "main/pkg/repository/interface"
	services "main/pkg/usecase/interface"
	"main/pkg/utils/models"
	"mime/multipart"
	"strconv"
)

type inventoryUseCase struct {
	repository interfaces.InventoryRepository
}

func NewInventoryUseCase(repo interfaces.InventoryRepository) services.InventoryUseCase {
	return &inventoryUseCase{
		repository: repo,
	}
}

func (i *inventoryUseCase) AddInventory(inventory models.Inventory, image *multipart.FileHeader) (models.InventoryResponse, error) {

	url, err := helper.AddImageToS3(image)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	inventory.Image = url
	//send the url and save it in database
	InventoryResponse, err := i.repository.AddInventory(inventory, url)
	if err != nil {
		fmt.Println(err)
		return models.InventoryResponse{}, err
	}

	return InventoryResponse, nil

}
func (i *inventoryUseCase) UpdateImage(invID int, image *multipart.FileHeader) (models.Inventory, error) {

	url, err := helper.AddImageToS3(image)
	if err != nil {
		return models.Inventory{}, err
	}
	//send the url and save it in database
	InventoryResponse, err := i.repository.UpdateImage(invID, url)
	if err != nil {
		fmt.Println(err)
		return models.Inventory{}, err
	}

	return InventoryResponse, nil

}
func (i *inventoryUseCase) UpdateInventory(invID int, invData models.UpdateInventory) (models.Inventory, error) {

	result, err := i.repository.CheckInventory(invID)
	if err != nil {

		return models.Inventory{}, err
	}

	if !result {
		return models.Inventory{}, errors.New("there is no inventory as you mentioned")
	}

	newinv, err := i.repository.UpdateInventory(invID, invData)
	if err != nil {
		return models.Inventory{}, err
	}

	return newinv, err
}

func (i *inventoryUseCase) DeleteInventory(inventoryID string) error {

	err := i.repository.DeleteInventory(inventoryID)
	if err != nil {
		return err
	}
	return nil

}

func (i *inventoryUseCase) ShowIndividualProducts(id string) (models.InventoryDetails, error) {

	product, err := i.repository.ShowIndividualProducts(id)
	if err != nil {
		return models.InventoryDetails{}, err
	}
	productID, err := strconv.Atoi(id)
	if err != nil {
		return models.InventoryDetails{}, err
	}
	var AdditionalImages []models.ImageInfo
	AdditionalImages, err = i.repository.GetImagesFromInventoryID(productID)
	if err != nil {
		return models.InventoryDetails{}, err
	}
	InvDetails := models.InventoryDetails{Inventory: product, AdditionalImages: AdditionalImages}

	return InvDetails, nil

}

func (i *inventoryUseCase) ListProducts(page int, limit int) ([]models.InventoryList, error) {

	productDetails, err := i.repository.ListProducts(page, limit)
	if err != nil {
		return []models.InventoryList{}, err
	}
	return productDetails, nil

}

// change return searchkey
func (i *inventoryUseCase) SearchProducts(key string, page, limit int) ([]models.InventoryList, error) {

	productDetails, err := i.repository.SearchProducts(key, page, limit)
	if err != nil {
		return []models.InventoryList{}, err
	}

	return productDetails, nil

}

func (i *inventoryUseCase) GetCategoryProducts(catID int, page, limit int) ([]models.InventoryList, error) {

	productDetails, err := i.repository.GetCategoryProducts(catID, page, limit)
	if err != nil {
		return []models.InventoryList{}, err
	}

	return productDetails, nil

}

func (i *inventoryUseCase) AddImage(product_id int, image *multipart.FileHeader) (models.InventoryResponse, error) {

	imageURL, err := helper.AddImageToS3(image)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	InventoryResponse, err := i.repository.AddImage(product_id, imageURL)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return InventoryResponse, nil

}

func (i *inventoryUseCase) DeleteImage(product_id int, image_id int) error {

	err := i.repository.DeleteImage(product_id, image_id)
	if err != nil {
		return err
	}

	return nil

}
