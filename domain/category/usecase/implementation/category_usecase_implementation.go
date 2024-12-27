package survey_category_usecase_implementation

import (
	survey_category_repository "github.com/celpung/gocleanarch/domain/category/repository"
	survey_category_usecase "github.com/celpung/gocleanarch/domain/category/usecase"
	"github.com/celpung/gocleanarch/entity"
)

type SurveyCategoryUsecaseStruct struct {
	repository survey_category_repository.SurveyCategoryRepository
}

// Create implements survey_category_usecase.SurveyCategoryUsecase.
func (u *SurveyCategoryUsecaseStruct) Create(category *entity.SurveyCategory) error {
	return u.repository.Create(category)
}

// Read implements survey_category_usecase.SurveyCategoryUsecase.
func (u *SurveyCategoryUsecaseStruct) Read(page int, limit int) ([]*entity.SurveyCategory, int64, error) {
	return u.repository.Read(page, limit)
}

// ReadById implements survey_category_usecase.SurveyCategoryUsecase.
func (u *SurveyCategoryUsecaseStruct) ReadById(id uint) (*entity.SurveyCategory, error) {
	return u.repository.ReadById(id)
}

// Update implements survey_category_usecase.SurveyCategoryUsecase.
func (u *SurveyCategoryUsecaseStruct) Update(category *entity.SurveyCategory) (*entity.SurveyCategory, error) {
	return u.repository.Update(category)
}

// Delete implements survey_category_usecase.SurveyCategoryUsecase.
func (u *SurveyCategoryUsecaseStruct) Delete(id uint) error {
	return u.repository.Delete(id)
}

func NewSurveyCategoryUsecase(repository survey_category_repository.SurveyCategoryRepository) survey_category_usecase.SurveyCategoryUsecase {
	return &SurveyCategoryUsecaseStruct{repository: repository}
}
