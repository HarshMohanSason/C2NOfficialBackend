package services

import (
	"c2nofficialsitebackend/database"
	"c2nofficialsitebackend/models"
	"errors"
)

func ProcessAddingCategoryData(category *models.Category) error {

	err := validateCategory(category)
	if err != nil {
		return err
	}
	categoryRepository := database.PostgresCategoryRepository{DB: database.GetDB()}
	err = categoryRepository.AddCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func validateCategory(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}
	if len(category.Name) > 50 {
		return errors.New("category name is too long")
	}
	return nil
}
