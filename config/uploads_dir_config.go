package config

import (
	"os"
)

func SetupUploadsDir() error {
	const uploadDirPath = "../uploads"
	if _, err := os.Stat(uploadDirPath); os.IsNotExist(err) {
		err := os.Mkdir(uploadDirPath, os.ModePerm)
		if err != nil {
			LogError(err)
		}
	}
	return nil
}
