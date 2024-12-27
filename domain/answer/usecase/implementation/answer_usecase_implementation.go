package survey_answer_usecase_impl

import (
	"errors"

	survey_answer_repository "github.com/celpung/gocleanarch/domain/answer/repository"
	survey_answer_usecase "github.com/celpung/gocleanarch/domain/answer/usecase"
	"github.com/celpung/gocleanarch/entity"
)

type surveyAnswerUsecaseImpl struct {
	repo survey_answer_repository.SurveyAnswerRepository
}

func NewSurveyAnswerUsecase(repo survey_answer_repository.SurveyAnswerRepository) survey_answer_usecase.SurveyAnswerUsecase {
	return &surveyAnswerUsecaseImpl{repo: repo}
}

func (u *surveyAnswerUsecaseImpl) Create(answer *entity.SurveyAnswer) error {
	if answer.Answer == "" {
		return errors.New("answer cannot be empty")
	}
	return u.repo.Create(answer)
}

func (u *surveyAnswerUsecaseImpl) Read(page, limit int) ([]entity.SurveyAnswer, int64, error) {
	return u.repo.Read(page, limit)
}

func (u *surveyAnswerUsecaseImpl) ReadById(id uint) (*entity.SurveyAnswer, error) {
	return u.repo.ReadById(id)
}

func (u *surveyAnswerUsecaseImpl) ReadByQuestionId(questionID uint, page, limit int) ([]entity.SurveyAnswer, int64, error) {
	return u.repo.ReadByQuestionId(questionID, page, limit)
}

func (u *surveyAnswerUsecaseImpl) Update(answer *entity.SurveyAnswer) (*entity.SurveyAnswer, error) {
	if answer.Answer == "" {
		return nil, errors.New("answer cannot be empty")
	}
	return u.repo.Update(answer)
}

func (u *surveyAnswerUsecaseImpl) Delete(id uint) error {
	return u.repo.Delete(id)
}
