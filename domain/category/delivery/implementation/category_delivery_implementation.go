package survey_category_delivery_implementation

import (
	"net/http"
	"strconv"

	survey_category_delivery "github.com/celpung/gocleanarch/domain/category/delivery"
	survey_category_usecase "github.com/celpung/gocleanarch/domain/category/usecase"
	"github.com/celpung/gocleanarch/entity"
	"github.com/gin-gonic/gin"
)

type CategoryUsecaseStruct struct {
	usecase survey_category_usecase.SurveyCategoryUsecase
}

// Create implements survey_category_delivery.SurveyCategoryDeliveryInterface.
func (d *CategoryUsecaseStruct) Create(c *gin.Context) {
	var reg entity.SurveyCategory

	// Bind the incoming JSON to the Survey struct
	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Call the usecase to create the survey
	if err := d.usecase.Create(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Return the created survey as the response
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    reg,
	})
}

// Read implements survey_category_delivery.SurveyCategoryDeliveryInterface.
func (d *CategoryUsecaseStruct) Read(c *gin.Context) {
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

	categories, total, err := d.usecase.Read(page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"categories": categories,
		"total_data": total,
	})
}

// ReadById implements survey_category_delivery.SurveyCategoryDeliveryInterface.
func (d *CategoryUsecaseStruct) ReadById(c *gin.Context) {
	//read the id from the url
	idStr := c.Param("id")

	//convert the id to uint
	id, err := strconv.ParseUint(idStr, 10, 64)

	//check if the id is valid
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid id parameter",
			"error":   err.Error(),
		})
		return
	}

	//call the usecase to get the survey by id
	category, err := d.usecase.ReadById(uint(id))

	//check if the survey is found
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	//return the survey as the response
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"category": category,
	})
}

// Update implements survey_category_delivery.SurveyCategoryDeliveryInterface.
func (d *CategoryUsecaseStruct) Update(c *gin.Context) {
	var reg entity.SurveyCategory

	// Bind the incoming JSON to the Survey struct
	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Call the usecase to update the survey
	cat, err := d.usecase.Update(&reg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Return the updated survey as the response
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"category": cat,
	})
}

// Delete implements survey_category_delivery.SurveyCategoryDeliveryInterface.
func (d *CategoryUsecaseStruct) Delete(c *gin.Context) {
	//read the id from the url
	idStr := c.Param("id")

	//convert the id to uint
	id, err := strconv.ParseUint(idStr, 10, 64)

	//check if the id is valid
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid id parameter",
			"error":   err.Error(),
		})
		return
	}

	//call the usecase to delete the survey
	err = d.usecase.Delete(uint(id))

	//check if the survey is found
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	//return the survey as the response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category deleted successfully",
	})
}

func NewSurveyCategoryDelivery(usecase survey_category_usecase.SurveyCategoryUsecase) survey_category_delivery.SurveyCategoryDeliveryInterface {
	return &CategoryUsecaseStruct{usecase: usecase}
}
