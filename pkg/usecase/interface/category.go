package interfaces

import (
	"main/pkg/domain"
)

type CategoryUseCase interface {
	AddCategory(category string) (domain.Category, error)
	UpdateCategory(current string, new string) (domain.Category, error)
	DeleteCategory(categoryID string) error
	GetCategories(page,limit int)([]domain.Category,error)
}
