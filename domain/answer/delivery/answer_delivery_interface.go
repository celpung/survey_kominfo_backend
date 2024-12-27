package survey_answer_delivery

import (
	"github.com/gin-gonic/gin"
)

type SurveyAnswerDelivery interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	ReadById(c *gin.Context)
	ReadByQuestionId(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
