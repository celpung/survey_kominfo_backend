package survey_usecase

import "github.com/celpung/gocleanarch/entity"

type SurveyUsecaseInterface interface {
	Create(survey *entity.Survey) (*entity.Survey, error)
	Read() ([]*entity.Survey, error)
	ReadByID(id uint) (*entity.Survey, error)
	Update(survey *entity.Survey) (*entity.Survey, error)
	Delete(id uint) error
}
