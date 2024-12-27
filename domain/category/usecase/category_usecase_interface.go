package survey_category_usecase

import "github.com/celpung/gocleanarch/entity"

type SurveyCategoryUsecase interface {
	Create(category *entity.SurveyCategory) error
	Read(page, limit int) ([]*entity.SurveyCategory, int64, error)
	ReadById(id uint) (*entity.SurveyCategory, error)
	Update(category *entity.SurveyCategory) (*entity.SurveyCategory, error)
	Delete(id uint) error
}
