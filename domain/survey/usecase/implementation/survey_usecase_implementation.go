package survey_usecase_implementation

import (
	survey_repository "github.com/celpung/gocleanarch/domain/survey/repository"
	survey_usecase "github.com/celpung/gocleanarch/domain/survey/usecase"
	"github.com/celpung/gocleanarch/entity"
)

type SurveyUsecaseStruct struct {
	repository survey_repository.SurveyRepositoryInterface
}

// Create implements survey_usecase.SurveyUsecaseInterface.
func (u *SurveyUsecaseStruct) Create(survey *entity.Survey) (*entity.Survey, error) {
	return u.repository.Create(survey)
}

// Read implements survey_usecase.SurveyUsecaseInterface.
func (u *SurveyUsecaseStruct) Read(page, limit int) ([]*entity.Survey, int64, error) {
	return u.repository.Read(page, limit)
}

func (u *SurveyUsecaseStruct) ReadByID(id uint) (*entity.Survey, error) {
	return u.repository.ReadByID(id)
}

// Update implements survey_usecase.SurveyUsecaseInterface.
func (u *SurveyUsecaseStruct) Update(survey *entity.Survey) (*entity.Survey, error) {
	// update survey
	return u.repository.Update(survey)
}

// Delete implements survey_usecase.SurveyUsecaseInterface.
func (u *SurveyUsecaseStruct) Delete(id uint) error {
	// delete
	return u.repository.Delete(id)
}

func NewSurveyUsecase(repository survey_repository.SurveyRepositoryInterface) survey_usecase.SurveyUsecaseInterface {
	return &SurveyUsecaseStruct{
		repository: repository,
	}
}
