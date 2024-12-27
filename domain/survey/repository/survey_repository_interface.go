package survey_repository

import "github.com/celpung/gocleanarch/entity"

type SurveyRepositoryInterface interface {
	Create(survey *entity.Survey) (*entity.Survey, error)
	Read() ([]*entity.Survey, error)
	ReadByID(id uint) (*entity.Survey, error)
	Update(survey *entity.Survey) (*entity.Survey, error)
	Delete(id uint) error
}
