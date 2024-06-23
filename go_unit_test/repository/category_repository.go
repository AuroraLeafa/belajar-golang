package repository

import "go_unit_test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
