package user_repository_implementation

import (
	"fmt"

	user_repository "github.com/celpung/gocleanarch/domain/user/repository"
	"github.com/celpung/gocleanarch/entity"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserRepositoryStruct struct {
	DB *gorm.DB
}

// Create implements user_repository.UserRepositoryInterface.
func (r *UserRepositoryStruct) Create(user *entity.User) (*entity.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		if gormErr, ok := err.(*mysql.MySQLError); ok {
			switch gormErr.Number {
			case 1062:
				return nil, fmt.Errorf("username '%s' sudah ada", user.Username)
			default:
				return nil, fmt.Errorf("database error: %v", gormErr.Message)
			}
		}
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return user, nil
}

// Read implements user_repository.UserRepositoryInterface.
// Read implements user_repository.UserRepositoryInterface.
func (r *UserRepositoryStruct) Read(page, limit int) ([]*entity.User, int64, error) {
	offset := (page - 1) * limit
	var users []*entity.User
	var totalCount int64

	// Count the total number of records
	if err := r.selectUserData(r.DB).Model(&entity.User{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated data
	if err := r.selectUserData(r.DB).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

// ReadByID implements user_repository.UserRepositoryInterface.
func (r *UserRepositoryStruct) ReadByID(userID uint) (*entity.User, error) {
	var user entity.User
	if err := r.selectUserData(r.DB).First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// ReadByEmail implements user_repository.UserRepositoryInterface.
func (r *UserRepositoryStruct) ReadByEmail(username string, isLogin bool) (*entity.User, error) {
	var user entity.User
	if isLogin {
		if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.selectUserData(r.DB).Where("username = ?", username).First(&user).Error; err != nil {
			return nil, err
		}
	}

	return &user, nil
}

// Update implements user_repository.UserRepositoryInterface.
func (r *UserRepositoryStruct) Update(user *entity.User) (*entity.User, error) {
	if err := r.DB.Model(&entity.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Delete implements user_repository.UserRepositoryInterface.
func (r *UserRepositoryStruct) Delete(userID uint) error {
	if err := r.DB.Delete(&entity.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryStruct) selectUserData(db *gorm.DB) *gorm.DB {
	return db.Select("ID, Name, Username, Active, Role")
}

func NewUserRepositry(db *gorm.DB) user_repository.UserRepositoryInterface {
	return &UserRepositoryStruct{DB: db}
}
