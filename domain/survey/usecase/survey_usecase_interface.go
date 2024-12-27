package survey_usecase

import "github.com/celpung/gocleanarch/entity"

type SurveyUsecaseInterface interface {
	Create(survey *entity.Survey) (*entity.Survey, error)
	Read(page, limit int) ([]*entity.Survey, int64, error)
	ReadByID(id uint) (*entity.Survey, error)
	ReadBySlug(slug string) (*entity.Survey, error)
	Update(survey *entity.Survey) (*entity.Survey, error)
	Delete(id uint) error
}
