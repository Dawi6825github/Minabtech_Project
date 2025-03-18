package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// UploadFile handles file uploads and saves them to a specified directory.
func UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	uploadDir := "uploads/"

	// Ensure the upload directory exists
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Generate a unique file name
	ext := filepath.Ext(fileHeader.Filename)
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, newFileName)

	// Create destination file
	destFile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %v", err)
	}
	defer destFile.Close()

	// Copy file content
	if _, err := io.Copy(destFile, file); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	// Return file path
	return filePath, nil
}
