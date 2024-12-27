package survey_delivery_implementation

import (
	"net/http"
	"strconv"

	survey_delivery "github.com/celpung/gocleanarch/domain/survey/delivery"
	survey_usecase "github.com/celpung/gocleanarch/domain/survey/usecase"
	"github.com/celpung/gocleanarch/entity"
	"github.com/gin-gonic/gin"
)

type SurveyDeliveryStruct struct {
	usecase survey_usecase.SurveyUsecaseInterface
}

// Create implements survey_delivery.SurveyDeliveryInterface.
func (d *SurveyDeliveryStruct) Create(c *gin.Context) {
	var reg entity.Survey
	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	survey, err := d.usecase.Create(&reg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    survey,
	})
}

// Read implements survey_delivery.SurveyDeliveryInterface.
func (d *SurveyDeliveryStruct) Read(c *gin.Context) {
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

	surveys, total, err := d.usecase.Read(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch surveys data!",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"message":    "Success fetch surveys data!",
		"surveys":    surveys,
		"total_data": total,
	})

	// surveys, total, err := d.usecase.Read(page, limit)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"success": false,
	// 		"message": "Failed to bind login data!",
	// 		"error":   err.Error(),
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"success":    true,
	// 	"message":    "Success fetch surveys data!",
	// 	"user":       surveys,
	// 	"total_data": total,
	// })
}

func (d *SurveyDeliveryStruct) ReadByID(c *gin.Context) {
	panic("unimplemented")
}

// Delete implements survey_delivery.SurveyDeliveryInterface.
func (d *SurveyDeliveryStruct) Delete(c *gin.Context) {
	panic("unimplemented")
}

// Update implements survey_delivery.SurveyDeliveryInterface.
func (d *SurveyDeliveryStruct) Update(c *gin.Context) {
	panic("unimplemented")
}

func NewSurveyDelivery(usecase survey_usecase.SurveyUsecaseInterface) survey_delivery.SurveyDeliveryInterface {
	return &SurveyDeliveryStruct{
		usecase: usecase,
	}
}
