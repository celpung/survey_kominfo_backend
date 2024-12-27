package survey_delivery

import "github.com/gin-gonic/gin"

type SurveyDeliveryInterface interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	ReadByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
