package user_delivery_implementation

import (
	"net/http"
	"strconv"

	user_delivery "github.com/celpung/gocleanarch/domain/user/delivery/gin"
	user_usecase "github.com/celpung/gocleanarch/domain/user/usecase"
	"github.com/celpung/gocleanarch/entity"
	"github.com/gin-gonic/gin"
)

type UserDeliveryStruct struct {
	UserUsecase user_usecase.UserUsecaseInterface
}

// Register implements user_delivery.UserDeliveryInterface.
func (d *UserDeliveryStruct) Register(c *gin.Context) {
	// get user input data
	var reg entity.User
	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed binding json data",
			"error":   err.Error(),
		})
		return
	}

	// perform registration
	user, err := d.UserUsecase.Create(&reg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to create user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Register success!",
		"user":    user,
	})
}

// Login implements user_delivery.UserDeliveryInterface.
func (d *UserDeliveryStruct) Login(c *gin.Context) {
	// get user input data
	type UserLogin struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var login UserLogin

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to bind login data!",
			"error":   err.Error(),
		})
		return
	}

	// perform login to get token data
	token, err := d.UserUsecase.Login(login.Username, login.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Login failed!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login success",
		"token":   token,
	})
}

// GetAllUserData implements user_delivery.UserDeliveryInterface.
func (d *UserDeliveryStruct) GetAllUserData(c *gin.Context) {
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

	user, total, err := d.UserUsecase.Read(page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to bind login data!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"message":    "Success fetch user data!",
		"user":       user,
		"total_data": total,
	})
}

func (d *UserDeliveryStruct) GetUserById(c *gin.Context) {
	// get user id from parameter
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid user id!",
			"error":   err.Error(),
		})
		return
	}

	user, err := d.UserUsecase.ReadByID(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to get user data!",
			"error":   err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success get user data!",
		"user":    user,
	})
}

func (d *UserDeliveryStruct) UpdateUser(c *gin.Context) {
	var updateData entity.UserUpdate
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to bind data!",
			"error":   err.Error(),
		})
		return
	}

	user, err := d.UserUsecase.Update(&updateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to update data!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success update user data!",
		"user":    user,
	})
}

func NewUserDelivery(usecase user_usecase.UserUsecaseInterface) user_delivery.UserDeliveryInterface {
	return &UserDeliveryStruct{
		UserUsecase: usecase,
	}
}
