package survey_question_usecase

import "github.com/celpung/gocleanarch/entity"

type SurveyQuestionUsecase interface {
	Create(question *entity.SurveyQuestion) error
	Read(page, limit int) ([]*entity.SurveyQuestion, int64, error)
	ReadById(id uint) (*entity.SurveyQuestion, error)
	ReadBySurveyId(surveyId uint, page, limit int) ([]*entity.SurveyQuestion, int64, error)
	Update(question *entity.SurveyQuestion) (*entity.SurveyQuestion, error)
	Delete(id uint) error
}
