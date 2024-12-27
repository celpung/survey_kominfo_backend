package uploader

import "github.com/gin-gonic/gin"

type FileUploader interface {
	UploadFile(c *gin.Context)
}
