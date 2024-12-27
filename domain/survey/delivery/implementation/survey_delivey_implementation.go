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

func (d *SurveyDeliveryStruct) Create(c *gin.Context) {
	userIDFloat := c.MustGet("userID").(float64)
	userID := int(userIDFloat)

	var reg entity.Survey

	// Set only the UserID, don't populate Author here
	reg.UserID = uint(userID)

	// Bind the incoming JSON to the Survey struct
	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Call the usecase to create the survey
	survey, err := d.usecase.Create(&reg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Return the created survey as the response
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
}

func (d *SurveyDeliveryStruct) ReadByID(c *gin.Context) {
	// get id from parameter
	idStr := c.Param("id")

	// convert id string to uint
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID parameter",
			"error":   err.Error(),
		})
		return
	}

	survey, err := d.usecase.ReadByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch survey data!",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success fetch survey data!",
		"survey":  survey,
	})
}

func (d *SurveyDeliveryStruct) ReadBySlug(c *gin.Context) {
	// get id from parameter
	slug := c.Param("slug")

	survey, err := d.usecase.ReadBySlug(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch survey data!",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success fetch survey data!",
		"survey":  survey,
	})
}

// Update implements survey_delivery.SurveyDeliveryInterface.
func (d *SurveyDeliveryStruct) Update(c *gin.Context) {
	userIDFloat := c.MustGet("userID").(float64)
	userID := int(userIDFloat)

	var reg entity.Survey

	reg.UserID = uint(userID)

	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	survey, err := d.usecase.Update(&reg)
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

// Delete implements survey_delivery.SurveyDeliveryInterface.
func (d *SurveyDeliveryStruct) Delete(c *gin.Context) {
	idStr := c.Param("id")

	// convert id string to uint
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID parameter",
			"error":   err.Error(),
		})
		return
	}

	if err = d.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete survey data!",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success delete survey data!",
	})
}

func NewSurveyDelivery(usecase survey_usecase.SurveyUsecaseInterface) survey_delivery.SurveyDeliveryInterface {
	return &SurveyDeliveryStruct{
		usecase: usecase,
	}
}
