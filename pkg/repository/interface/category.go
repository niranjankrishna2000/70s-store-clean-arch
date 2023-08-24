package interfaces

import "main/pkg/domain"

type CategoryRepository interface {
	AddCategory(category string) (domain.Category, error)
	CheckCategory(currrent string) (bool, error)
	UpdateCategory(current, new string) (domain.Category, error)
	DeleteCategory(categoryID string) error
}
