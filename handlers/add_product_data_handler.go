package handlers

import (
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/services"
	"c2nofficialsitebackend/utils"
	"fmt"
	"net/http"
	"strconv"
)

func AddProductData(response http.ResponseWriter, receivedRequest *http.Request) {
	if receivedRequest.Method != http.MethodPost {
		http.Error(response, "Error occurred, wrong type of request", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form
	err := receivedRequest.ParseMultipartForm(10 << 20) // 10MB limit for the images in total.
	if err != nil {
		http.Error(response, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	// Validate required fields
	productName := receivedRequest.FormValue("name")
	if productName == "" {
		http.Error(response, "Product name is required", http.StatusBadRequest)
		return
	}
	categoryIDStr := receivedRequest.FormValue("category_id")

	if categoryIDStr == "" {
		http.Error(response, "A Category is required", http.StatusBadRequest)
		return
	}
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		http.Error(response, fmt.Sprintf("Invalid category ID: %v", err), http.StatusBadRequest)
		return
	}

	priceStr := receivedRequest.FormValue("price")
	priceUint64, err := strconv.ParseUint(priceStr, 10, 32)
	if err != nil {
		http.Error(response, "Invalid price value", http.StatusBadRequest)
		return
	}

	discountStr := receivedRequest.FormValue("discount")
	discountUint64, err := strconv.ParseUint(discountStr, 10, 32)
	if err != nil {
		http.Error(response, "Invalid discount value", http.StatusBadRequest)
		return
	}

	// Extract form fields for product creation
	inventoryStr := receivedRequest.FormValue("inventory")
	inventoryUint64, err := strconv.ParseUint(inventoryStr, 10, 16)
	if err != nil {
		http.Error(response, "Invalid inventory value", http.StatusBadRequest)
		return
	}

	statusStr := receivedRequest.FormValue("status")
	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		http.Error(response, "Invalid status value", http.StatusBadRequest)
		return
	}
	weight, err := strconv.ParseFloat(receivedRequest.FormValue("weight"), 64)
	if err != nil {
		http.Error(response, "Invalid weight value", http.StatusBadRequest)
		return
	}
	length, err := strconv.ParseFloat(receivedRequest.FormValue("length"), 64)
	if err != nil {
		http.Error(response, "Invalid length value", http.StatusBadRequest)
		return
	}
	width, err := strconv.ParseFloat(receivedRequest.FormValue("width"), 64)
	if err != nil {
		http.Error(response, "Invalid width value", http.StatusBadRequest)
		return
	}
	height, err := strconv.ParseFloat(receivedRequest.FormValue("height"), 64)
	if err != nil {
		http.Error(response, "Invalid height value", http.StatusBadRequest)
	}
	// Upload thumbnail image
	_, thumbnailFileHeader, err := receivedRequest.FormFile("thumbnail_image")
	if err != nil {
		http.Error(response, "Unable to parse thumbnail image, Bad Request", http.StatusBadRequest)
		return
	}
	thumbnailImage, err := utils.UploadSingleFile(thumbnailFileHeader, productName, "../uploads/products")
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	// Upload carousel images
	carouselFileHeaders := receivedRequest.MultipartForm.File["carousel_images"]
	if carouselFileHeaders == nil {
		http.Error(response, "No carousel images found", http.StatusBadRequest)
		return
	}
	var carouselImages []string
	carouselImages, err = utils.UploadMultipleFiles(carouselFileHeaders, productName, "../uploads/products")
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	// Create the product
	product := &models.Product{
		Name:             productName,
		CategoryID:       uint(categoryID),
		LongDescription:  receivedRequest.FormValue("long_description"),
		ShortDescription: receivedRequest.FormValue("short_description"),
		ThumbnailImage:   thumbnailImage,
		CarouselImages:   carouselImages,
		Slug:             receivedRequest.FormValue("slug"),
		Price:            uint32(priceUint64),
		Discount:         uint32(discountUint64),
		Inventory:        uint16(inventoryUint64),
		SKU:              receivedRequest.FormValue("sku"),
		Status:           status,
		Weight:           weight,
		Length:           length,
		Width:            width,
		Height:           height,
	}
	fmt.Printf("%+v\n", product)
	err = services.ProcessAddingProductData(product)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

}
