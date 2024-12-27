package survey_repository_implementation

import (
	survey_repository "github.com/celpung/gocleanarch/domain/survey/repository"
	"github.com/celpung/gocleanarch/entity"
	"gorm.io/gorm"
)

type SurveyRepositoryStruct struct {
	DB *gorm.DB
}

// Create implements survey_repository.SurveyRepositoryInterface.
func (r *SurveyRepositoryStruct) Create(survey *entity.Survey) (*entity.Survey, error) {
	// insert data into databse
	if err := r.DB.Create(survey).Error; err != nil {
		return nil, err
	}

	return survey, nil
}

// Read implements survey_repository.SurveyRepositoryInterface.
func (r *SurveyRepositoryStruct) Read(page, limit int) ([]*entity.Survey, int64, error) {
	// get all data from database
	offset := (page - 1) * limit
	var totalCount int64
	var surveys []*entity.Survey

	if err := r.DB.Model(&entity.Survey{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Select specific fields for the User model, excluding the Password field
	if err := r.DB.Limit(limit).Offset(offset).
		Preload("Questions").
		Preload("Category", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, created_at, updated_at, deleted_at")
		}).
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, username, active, role, created_at, updated_at, deleted_at")
		}).
		Find(&surveys).Error; err != nil {
		return nil, 0, err
	}

	return surveys, totalCount, nil
}

// ReadByID implements survey_repository.SurveyRepositoryInterface.
func (r *SurveyRepositoryStruct) ReadByID(id uint) (*entity.Survey, error) {
	// read survey by id
	var survey entity.Survey
	if err := r.DB.Preload("Questions").
		Preload("Category", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, created_at, updated_at, deleted_at")
		}).
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, username, active, role, created_at, updated_at, deleted_at")
		}).First(&survey, id).Error; err != nil {
		return nil, err
	}

	return &survey, nil
}

func (r *SurveyRepositoryStruct) ReadBySlug(slug string) (*entity.Survey, error) {
	// read survey by slug
	var survey entity.Survey
	if err := r.DB.Preload("Questions").
		Preload("Category", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, created_at, updated_at, deleted_at")
		}).
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, username, active, role, created_at, updated_at, deleted_at")
		}).Where("slug = ?", slug).First(&survey).Error; err != nil {
		return nil, err
	}

	return &survey, nil
}

// Update implements survey_repository.SurveyRepositoryInterface.
func (r *SurveyRepositoryStruct) Update(survey *entity.Survey) (*entity.Survey, error) {
	// update the survey by id
	if err := r.DB.Model(&entity.Survey{}).Where("id = ?", survey.ID).Updates(survey).Error; err != nil {
		return nil, err
	}

	return survey, nil
}

// Delete implements survey_repository.SurveyRepositoryInterface.
func (r *SurveyRepositoryStruct) Delete(id uint) error {
	// delete by id
	if err := r.DB.Delete(&entity.Survey{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewSurveyRepositry(db *gorm.DB) survey_repository.SurveyRepositoryInterface {
	return &SurveyRepositoryStruct{DB: db}
}
