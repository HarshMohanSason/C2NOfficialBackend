package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
)

func UploadSingleFile(fileHeader *multipart.FileHeader, productName string, dirPath string) (string, error) {
	if !DoesDirExists(dirPath) {
		err := CreateDir(dirPath)
		if err != nil {
			return "", err
		}
	}

	getExtension := GetFileExtensionFromFileHeader(fileHeader)
	err := IsAllowedExtension(getExtension)
	if err != nil {
		return "", err
	}
	filePath := GenerateUniqueFilePath(productName, getExtension)

	filePath = filepath.Join(dirPath, filePath)

	newFile, err := CreateFile(filePath)

	//Open the multipart file from the header
	file, err := fileHeader.Open()
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to open the file: %v", err))
	}
	defer file.Close()

	err = CopyFileContentsFromMultiPartFile(newFile, &file)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func UploadMultipleFiles(fileHeaders []*multipart.FileHeader, productName string, dirPath string) ([]string, error) {
	var uploadedPaths []string
	if !DoesDirExists(dirPath) {
		err := CreateDir(dirPath)
		if err != nil {
			return uploadedPaths, err
		}
	}
	for _, fileHeader := range fileHeaders {
		getExtension := GetFileExtensionFromFileHeader(fileHeader)
		err := IsAllowedExtension(getExtension)
		if err != nil {
			return nil, err
		}
		filePath := GenerateUniqueFilePath(productName, getExtension)
		filePath = filepath.Join(dirPath, filePath)
		newFile, err := CreateFile(filePath)

		//Open the multipart file from the header
		file, err := fileHeader.Open()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to open the file: %v", err))
		}

		err = CopyFileContentsFromMultiPartFile(newFile, &file)
		if err != nil {
			return nil, err
		}
		file.Close()
		uploadedPaths = append(uploadedPaths, filePath)
	}
	return uploadedPaths, nil
}
