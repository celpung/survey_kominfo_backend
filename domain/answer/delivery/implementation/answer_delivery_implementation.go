package survey_answer_delivery_impl

import (
	"net/http"
	"strconv"

	survey_answer_delivery "github.com/celpung/gocleanarch/domain/answer/delivery"
	survey_answer_usecase "github.com/celpung/gocleanarch/domain/answer/usecase"
	"github.com/celpung/gocleanarch/entity"
	"github.com/gin-gonic/gin"
)

type surveyAnswerDeliveryImpl struct {
	usecase survey_answer_usecase.SurveyAnswerUsecase
}

func NewSurveyAnswerDelivery(usecase survey_answer_usecase.SurveyAnswerUsecase) survey_answer_delivery.SurveyAnswerDelivery {
	return &surveyAnswerDeliveryImpl{usecase: usecase}
}

func (d *surveyAnswerDeliveryImpl) Create(c *gin.Context) {
	var answer entity.SurveyAnswer

	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := d.usecase.Create(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": answer})
}

func (d *surveyAnswerDeliveryImpl) Read(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	answers, total, err := d.usecase.Read(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":   total,
		"answers": answers,
	})
}

func (d *surveyAnswerDeliveryImpl) ReadById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	answer, err := d.usecase.ReadById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": answer})
}

func (d *surveyAnswerDeliveryImpl) ReadByQuestionId(c *gin.Context) {
	questionIDStr := c.Param("question_id")
	questionID, err := strconv.ParseUint(questionIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Question ID"})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	answers, total, err := d.usecase.ReadByQuestionId(uint(questionID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"succes":     true,
		"total_data": total,
		"answers":    answers,
	})
}

func (d *surveyAnswerDeliveryImpl) Update(c *gin.Context) {
	var answer entity.SurveyAnswer

	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAnswer, err := d.usecase.Update(&answer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedAnswer})
}

func (d *surveyAnswerDeliveryImpl) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := d.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Answer deleted successfully"})
}
