package survey_question_delivery_implementation

import (
	"net/http"
	"strconv"

	survey_question_delivery "github.com/celpung/gocleanarch/domain/question/delivery"
	survey_question_usecase "github.com/celpung/gocleanarch/domain/question/usecase"
	"github.com/celpung/gocleanarch/entity"
	"github.com/gin-gonic/gin"
)

type QuestionDeliveryStruct struct {
	usecase survey_question_usecase.SurveyQuestionUsecase
}

// Create implements survey_question_delivery.SurveyQuestionDeliveryInterface.
func (d *QuestionDeliveryStruct) Create(c *gin.Context) {
	var question entity.SurveyQuestion

	// Bind JSON to the SurveyQuestion struct
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Call usecase to create the question
	if err := d.usecase.Create(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    question,
	})
}

// Read implements survey_question_delivery.SurveyQuestionDeliveryInterface.
func (d *QuestionDeliveryStruct) Read(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid page parameter",
			"error":   err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid limit parameter",
			"error":   err.Error(),
		})
		return
	}

	questions, total, err := d.usecase.Read(page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"questions":  questions,
		"total_data": total,
	})
}

// ReadById implements survey_question_delivery.SurveyQuestionDeliveryInterface.
func (d *QuestionDeliveryStruct) ReadById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid id parameter",
			"error":   err.Error(),
		})
		return
	}

	question, err := d.usecase.ReadById(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"question": question,
	})
}

func (d *QuestionDeliveryStruct) ReadBySurveyId(c *gin.Context) {
	surveyIdStr := c.Param("survey_id")

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid page parameter",
			"error":   err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid limit parameter",
			"error":   err.Error(),
		})
		return
	}

	surveyId, err := strconv.ParseUint(surveyIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid survey_id parameter",
			"error":   err.Error(),
		})
		return
	}

	questions, total, err := d.usecase.ReadBySurveyId(uint(surveyId), page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"questions": questions,
		"total_data":     total,
	})
}

// Update implements survey_question_delivery.SurveyQuestionDeliveryInterface.
func (d *QuestionDeliveryStruct) Update(c *gin.Context) {
	var question entity.SurveyQuestion

	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	updatedQuestion, err := d.usecase.Update(&question)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"question": updatedQuestion,
	})
}

// Delete implements survey_question_delivery.SurveyQuestionDeliveryInterface.
func (d *QuestionDeliveryStruct) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid id parameter",
			"error":   err.Error(),
		})
		return
	}

	err = d.usecase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Question deleted successfully",
	})
}

// NewSurveyQuestionDelivery creates a new instance of QuestionDeliveryStruct.
func NewSurveyQuestionDelivery(usecase survey_question_usecase.SurveyQuestionUsecase) survey_question_delivery.SurveyQuestionDeliveryInterface {
	return &QuestionDeliveryStruct{usecase: usecase}
}
