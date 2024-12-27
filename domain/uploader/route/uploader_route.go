package uploader_router

import (
	uploader_implementation "github.com/celpung/gocleanarch/domain/uploader/implementation"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	fileUploader := uploader_implementation.NewFileUploaderImplementation()

	// Define the routes
	routes := r.Group("/upload")
	{
		routes.POST("", fileUploader.UploadFile)
	}
}
