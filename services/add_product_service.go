package services

import (
	"c2nofficialsitebackend/database"
	"c2nofficialsitebackend/models"
	"errors"
	"path/filepath"
	"strings"
	"time"
)

func ProcessAddingProductData(product *models.Product) error {
	err := validateProductData(product)
	if err != nil {
		return err
	}
	productRepository := database.PostgresProductRepository{DB: database.GetDB()}
	err = productRepository.AddProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func validateProductData(product *models.Product) error {

	err := validateName(product)
	if err != nil {
		return err
	}
	if product.LongDescription == "" {
		return errors.New("product long_description cannot be empty")
	}
	if product.ShortDescription == "" {
		return errors.New("product short_description cannot be empty")
	}
	err = validateImages(product)
	if err != nil {
		return err
	}
	err = validateSlug(product)
	if err != nil {
		return err
	}
	err = validatePrice(product)
	if err != nil {
		return err
	}
	err = validateProductWeightAndDimensions(product)
	if err != nil {
		return err
	}
	return nil
}

func validateName(product *models.Product) error {
	if product.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if len(product.Name) > 100 {
		return errors.New("product name cannot be longer than 100 characters")
	}
	if len(product.Name) < 10 {
		return errors.New("product name cannot be shorter than 10 characters")
	}
	return nil
}

func validateImages(product *models.Product) error {
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}
	if product.ThumbnailImage == "" {
		return errors.New("product thumbnail_image cannot be empty")
	}
	if product.CarouselImages == nil {
		return errors.New("product carousel_images cannot be empty")
	}
	ext := filepath.Ext(product.ThumbnailImage)
	if !allowedExtensions[ext] {
		return errors.New("product thumbnail_image must be one of jpg,jpeg,png")
	}
	for _, image := range product.CarouselImages {
		imageExtension := filepath.Ext(image)
		if !allowedExtensions[imageExtension] {
			return errors.New("product thumbnail_image must be one of jpg,jpeg,png")
		}
	}
	return nil
}

func validateSlug(product *models.Product) error {
	if product.Slug == "" {
		return errors.New("product slug cannot be empty")
	}
	//Unique slug with the current timestamp
	product.Slug = strings.ToLower(product.Slug) + time.Now().String()
	return nil
}

func validatePrice(product *models.Product) error {
	if product.Price == 0 {
		return errors.New("product price cannot be zero")
	}
	if product.Price > 10000 {
		return errors.New("product price cannot be greater than 10000 rupees")
	}
	if product.Price < 100 {
		return errors.New("product price cannot be less than 100 rupees")
	}
	return nil
}

func validateProductWeightAndDimensions(product *models.Product) error {
	if product.Weight == 0 {
		return errors.New("product weight cannot be zero")
	}
	if product.Weight > 1000 {
		return errors.New("product weight cannot be greater than 1000g")
	}
	if product.Length == 0 {
		return errors.New("product length cannot be zero")
	}
	if product.Width == 0 {
		return errors.New("product width cannot be zero")
	}
	if product.Height == 0 {
		return errors.New("product height cannot be zero")
	}
	return nil
}
