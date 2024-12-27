package uploader_implementation

import (
	"fmt"
	"net/http"

	"github.com/celpung/gocleanarch/configs/environment"
	"github.com/celpung/gocleanarch/domain/uploader"
	"github.com/celpung/gouploader"
	"github.com/gin-gonic/gin"
)

type FileUploaderImplementationStruct struct{}

// UploadFile implements uploader.FileUploader.
func (f *FileUploaderImplementationStruct) UploadFile(c *gin.Context) {
	// Handle multiple file uploads
	uploadedFiles, err := gouploader.Multiple(c.Request, "./public/files", "file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to upload files",
			"error":   err.Error(),
		})
		return
	}

	var fileUrls []string

	// If files are uploaded successfully, generate URLs for each file
	if len(uploadedFiles) > 0 {
		for _, file := range uploadedFiles {
			fileUrl := fmt.Sprintf("%s/files/%s", environment.Env.BASE_URL, file.Filename)
			fileUrls = append(fileUrls, fileUrl)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Files uploaded successfully",
		"urls":    fileUrls, // Return the array of file URLs
	})
}

func NewFileUploaderImplementation() uploader.FileUploader {
	return &FileUploaderImplementationStruct{}
}
