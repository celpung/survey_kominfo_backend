package survey_category_delivery

import (
	"github.com/gin-gonic/gin"
)

type SurveyCategoryDeliveryInterface interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	ReadById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
