package survey_answer_usecase

import (
	"github.com/celpung/gocleanarch/entity"
)

type SurveyAnswerUsecase interface {
	Create(answer *entity.SurveyAnswer) error
	Read(page, limit int) ([]entity.SurveyAnswer, int64, error)
	ReadById(id uint) (*entity.SurveyAnswer, error)
	ReadByQuestionId(questionID uint, page, limit int) ([]entity.SurveyAnswer, int64, error)
	Update(answer *entity.SurveyAnswer) (*entity.SurveyAnswer, error)
	Delete(id uint) error
}
