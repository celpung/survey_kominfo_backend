package survey_answer_repository_impl

import (
	survey_answer_repository "github.com/celpung/gocleanarch/domain/answer/repository"
	"github.com/celpung/gocleanarch/entity"
	"gorm.io/gorm"
)

type surveyAnswerRepositoryImpl struct {
	db *gorm.DB
}

func NewSurveyAnswerRepository(db *gorm.DB) survey_answer_repository.SurveyAnswerRepository {
	return &surveyAnswerRepositoryImpl{db: db}
}

func (r *surveyAnswerRepositoryImpl) Create(answer *entity.SurveyAnswer) error {
	return r.db.Create(answer).Error
}

func (r *surveyAnswerRepositoryImpl) Read(page, limit int) ([]entity.SurveyAnswer, int64, error) {
	var answers []entity.SurveyAnswer
	var total int64

	offset := (page - 1) * limit
	err := r.db.Model(&entity.SurveyAnswer{}).Count(&total).Limit(limit).Offset(offset).Find(&answers).Error
	return answers, total, err
}

func (r *surveyAnswerRepositoryImpl) ReadById(id uint) (*entity.SurveyAnswer, error) {
	var answer entity.SurveyAnswer
	err := r.db.First(&answer, id).Error
	return &answer, err
}

func (r *surveyAnswerRepositoryImpl) ReadByQuestionId(questionID uint, page, limit int) ([]entity.SurveyAnswer, int64, error) {
	var answers []entity.SurveyAnswer
	var total int64

	offset := (page - 1) * limit
	err := r.db.Model(&entity.SurveyAnswer{}).Where("question_id = ?", questionID).Count(&total).Limit(limit).Offset(offset).Find(&answers).Error
	return answers, total, err
}

func (r *surveyAnswerRepositoryImpl) Update(answer *entity.SurveyAnswer) (*entity.SurveyAnswer, error) {
	err := r.db.Save(answer).Error
	return answer, err
}

func (r *surveyAnswerRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entity.SurveyAnswer{}, id).Error
}
