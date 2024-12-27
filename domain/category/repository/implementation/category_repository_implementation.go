package survey_category_repository_implementation

import (
	survey_category_repository "github.com/celpung/gocleanarch/domain/category/repository"
	"github.com/celpung/gocleanarch/entity"
	"gorm.io/gorm"
)

type SurveyCategoryRepositoryStruct struct {
	DB *gorm.DB
}

// Create implements survey_category_repository.SurveyCategoryRepository.
func (r *SurveyCategoryRepositoryStruct) Create(category *entity.SurveyCategory) error {
	return r.DB.Create(category).Error
}

// Read implements survey_category_repository.SurveyCategoryRepository.
func (r *SurveyCategoryRepositoryStruct) Read(page, limit int) ([]*entity.SurveyCategory, int64, error) {
	offset := (page - 1) * limit
	var totalCount int64
	var categories []*entity.SurveyCategory

	if err := r.DB.Model(&entity.SurveyCategory{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Limit(limit).Offset(offset).
		Preload("Surveys").
		Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, totalCount, nil
}

// ReadById implements survey_category_repository.SurveyCategoryRepository.
func (r *SurveyCategoryRepositoryStruct) ReadById(id uint) (*entity.SurveyCategory, error) {
	var category entity.SurveyCategory
	if err := r.DB.Preload("Surveys").First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// Update implements survey_category_repository.SurveyCategoryRepository.
func (r *SurveyCategoryRepositoryStruct) Update(category *entity.SurveyCategory) (*entity.SurveyCategory, error) {
	// update the category
	if err := r.DB.Model(&entity.SurveyCategory{}).Where("id = ?", category.ID).Updates(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

// Delete implements survey_category_repository.SurveyCategoryRepository.
func (r *SurveyCategoryRepositoryStruct) Delete(id uint) error {
	// delete the category
	if err := r.DB.Where("id = ?", id).Delete(&entity.SurveyCategory{}).Error; err != nil {
		//return error
		return err
	}

	return nil
}

// NewSurveyCategoryRepository creates a new instance of surveyCategoryRepository.
func NewSurveyCategoryRepository(db *gorm.DB) survey_category_repository.SurveyCategoryRepository {
	return &SurveyCategoryRepositoryStruct{db}
}
