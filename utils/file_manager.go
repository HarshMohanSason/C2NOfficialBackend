package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DoesDirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func CreateDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func GetFileExtensionFromFileHeader(fileHeader *multipart.FileHeader) string {
	return filepath.Ext(fileHeader.Filename)
}

// GenerateUniqueFilePath with the current time along with the file extension.
func GenerateUniqueFilePath(customName string, ext string) string {
	customName = strings.TrimSpace(customName)
	customName = strings.ReplaceAll(customName, " ", "_")
	customName = strings.ToLower(customName)
	return customName + time.Now().Format("20060102_150405.000000000") + ext
}

func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func CopyFileContentsFromMultiPartFile(destinationFile *os.File, originalFile *multipart.File) error {
	// Copy the uploaded file into the newly created file
	_, err := io.Copy(destinationFile, *originalFile)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to copy file contents: %v", err))
	}
	destinationFile.Close() //close the file
	return nil
}

func IsAllowedExtension(ext string) error {
	availableExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
		".pdf":  true,
	}
	if _, ok := availableExtensions[ext]; !ok {
		return errors.New("unsupported file extension")
	}
	return nil
}
