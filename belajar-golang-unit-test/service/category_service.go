package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service *CategoryService) Get(Id string) (*entity.Category, error) {
	result := service.Repository.FindById(Id)
	if result == nil {
		return nil, errors.New("Category not found")
	}
	return result, nil
}
