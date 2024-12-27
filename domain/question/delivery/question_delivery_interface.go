package survey_question_delivery

import (
	"github.com/gin-gonic/gin"
)

type SurveyQuestionDeliveryInterface interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	ReadById(c *gin.Context)
	ReadBySurveyId(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
