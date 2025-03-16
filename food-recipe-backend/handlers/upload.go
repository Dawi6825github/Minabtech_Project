package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"your_project_path/golang/utils"
)

// UploadFile handles file uploads
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the form data to allow file uploads
	err := r.ParseMultipartForm(10 << 20) // Limit file size to 10MB
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Unable to parse form data")
		return
	}

	// Retrieve the file from the form
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "No file provided")
		return
	}
	defer file.Close()

	// Set the upload destination folder
	uploadDir := "./uploads"
	// Ensure the uploads directory exists
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Unable to create upload directory")
		return
	}

	// Generate a unique file name for the uploaded file
	fileName := fmt.Sprintf("%s/%s", uploadDir, "uploaded_file.jpg")
	dst, err := os.Create(fileName)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Unable to save file")
		return
	}
	defer dst.Close()

	// Copy the content of the uploaded file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Unable to save file content")
		return
	}

	// Respond with the file path or file URL if uploaded successfully
	fileURL := fmt.Sprintf("/uploads/%s", filepath.Base(fileName))
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"file_url": fileURL})
}

// ServeFile serves uploaded files to the client
func ServeFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileName := params["filename"]

	// Set the path for the file
	filePath := fmt.Sprintf("./uploads/%s", fileName)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		utils.RespondWithError(w, http.StatusNotFound, "File not found")
		return
	}

	// Serve the file
	http.ServeFile(w, r, filePath)
}
