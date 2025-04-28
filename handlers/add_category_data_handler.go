package handlers

import (
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/services"
	"c2nofficialsitebackend/utils"
	"net/http"
)

func AddCategoryData(response http.ResponseWriter, receivedRequest *http.Request) {

	if receivedRequest.Method != http.MethodPost {
		http.Error(response, "Wrong method", http.StatusMethodNotAllowed)
		return
	}
	err := receivedRequest.ParseMultipartForm(1 << 20) // 1MB limit for the size_chart
	if err != nil {
		http.Error(response, "Unable to parse form data", http.StatusBadRequest)
		return
	}
	categoryName := receivedRequest.FormValue("name")
	if categoryName == "" {
		http.Error(response, "Category name is required", http.StatusBadRequest)
		return
	}
	sizeChart := receivedRequest.FormValue("size_chart")
	if sizeChart == "" {
		http.Error(response, "Size chart is required", http.StatusBadRequest)
		return
	}
	_, howToMeasureImageHeader, err := receivedRequest.FormFile("how_to_measure_image")
	if err != nil {
		http.Error(response, "Unable to parse how to measure image, Bad Request", http.StatusBadRequest)
		return
	}
	howToMeasureImage, err := utils.UploadSingleFile(howToMeasureImageHeader, categoryName, "../uploads/categories")
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	var customizationPDF *string
	_, customizationPDFHeader, err := receivedRequest.FormFile("customization_pdf")
	//customizationPDF can be nil. So if it is storing nil value in the category table as is.
	if customizationPDFHeader == nil {
		customizationPDF = nil
	} else { //If not nil, proceed to upload the pdf file.
		if err != nil {
			http.Error(response, "Unable to customization pdf, bad request", http.StatusBadRequest)
			return
		}
		*customizationPDF, err = utils.UploadSingleFile(customizationPDFHeader, categoryName, "../uploads/categories")
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}
	}

	category := &models.Category{
		Name:              categoryName,
		SizeChart:         sizeChart,
		HowToMeasureImage: howToMeasureImage,
		CustomizationPDF:  customizationPDF,
	}

	err = services.ProcessAddingCategoryData(category)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
}
