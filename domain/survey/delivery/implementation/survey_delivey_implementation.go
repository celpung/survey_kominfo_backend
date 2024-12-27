package survey_delivery_implementation

import (
	"net/http"

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

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    survey,
	})
}

// Read implements survey_delivery.SurveyDeliveryInterface.
func (d *SurveyDeliveryStruct) Read(c *gin.Context) {
	panic("unimplemented")
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
