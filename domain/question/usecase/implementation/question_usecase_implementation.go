package survey_question_usecase_implementation

import (
	survey_question_repository "github.com/celpung/gocleanarch/domain/question/repository"
	survey_question_usecase "github.com/celpung/gocleanarch/domain/question/usecase"
	"github.com/celpung/gocleanarch/entity"
)

type SurveyQuestionUsecaseStruct struct {
	repository survey_question_repository.SurveyQuestionRepository
}

// Create implements survey_question_usecase.SurveyQuestionUsecase.
func (u *SurveyQuestionUsecaseStruct) Create(question *entity.SurveyQuestion) error {
	return u.repository.Create(question)
}

// Read implements survey_question_usecase.SurveyQuestionUsecase.
func (u *SurveyQuestionUsecaseStruct) Read(page, limit int) ([]*entity.SurveyQuestion, int64, error) {
	return u.repository.Read(page, limit)
}

// ReadById implements survey_question_usecase.SurveyQuestionUsecase.
func (u *SurveyQuestionUsecaseStruct) ReadById(id uint) (*entity.SurveyQuestion, error) {
	return u.repository.ReadById(id)
}

func (u *SurveyQuestionUsecaseStruct) ReadBySurveyId(surveyId uint, page, limit int) ([]*entity.SurveyQuestion, int64, error) {
	return u.repository.ReadBySurveyId(surveyId, page, limit)
}

// Update implements survey_question_usecase.SurveyQuestionUsecase.
func (u *SurveyQuestionUsecaseStruct) Update(question *entity.SurveyQuestion) (*entity.SurveyQuestion, error) {
	return u.repository.Update(question)
}

// Delete implements survey_question_usecase.SurveyQuestionUsecase.
func (u *SurveyQuestionUsecaseStruct) Delete(id uint) error {
	return u.repository.Delete(id)
}

// NewSurveyQuestionUsecase creates a new instance of SurveyQuestionUsecase.
func NewSurveyQuestionUsecase(repository survey_question_repository.SurveyQuestionRepository) survey_question_usecase.SurveyQuestionUsecase {
	return &SurveyQuestionUsecaseStruct{repository: repository}
}
