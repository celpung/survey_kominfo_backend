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
	var paths []string
	var originalNames []string

	// If files are uploaded successfully, generate URLs for each file
	if len(uploadedFiles) > 0 {
		for _, file := range uploadedFiles {
			fileUrl := fmt.Sprintf("%s/files/%s", environment.Env.BASE_URL, file.Filename)
			fileUrls = append(fileUrls, fileUrl)

			filePath := fmt.Sprintf("/files/%s", file.Filename)
			paths = append(paths, filePath)

			fileHeader, _ := c.FormFile("file") // You can use the field name to get FileHeader
			if fileHeader != nil {
				originalNames = append(originalNames, fileHeader.Filename)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success":        true,
		"message":        "Files uploaded successfully",
		"urls":           fileUrls,
		"paths":          paths,
		"original_names": originalNames,
	})
}

func NewFileUploaderImplementation() uploader.FileUploader {
	return &FileUploaderImplementationStruct{}
}
