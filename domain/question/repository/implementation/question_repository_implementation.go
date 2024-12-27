package survey_question_repository_implementation

import (
	survey_question_repository "github.com/celpung/gocleanarch/domain/question/repository"
	"github.com/celpung/gocleanarch/entity"
	"gorm.io/gorm"
)

type SurveyQuestionRepositoryStruct struct {
	DB *gorm.DB
}

// Create implements survey_question_repository.SurveyQuestionRepository.
func (r *SurveyQuestionRepositoryStruct) Create(question *entity.SurveyQuestion) error {
	return r.DB.Create(question).Error
}

// Read implements survey_question_repository.SurveyQuestionRepository.
func (r *SurveyQuestionRepositoryStruct) Read(page, limit int) ([]*entity.SurveyQuestion, int64, error) {
	offset := (page - 1) * limit
	var totalCount int64
	var questions []*entity.SurveyQuestion

	if err := r.DB.Model(&entity.SurveyQuestion{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Limit(limit).Offset(offset).
		Preload("Survey").
		Find(&questions).Error; err != nil {
		return nil, 0, err
	}

	return questions, totalCount, nil
}

// ReadById implements survey_question_repository.SurveyQuestionRepository.
func (r *SurveyQuestionRepositoryStruct) ReadById(id uint) (*entity.SurveyQuestion, error) {
	var question entity.SurveyQuestion
	if err := r.DB.Preload("Survey").First(&question, id).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *SurveyQuestionRepositoryStruct) ReadBySurveyId(surveyId uint, page, limit int) ([]*entity.SurveyQuestion, int64, error) {
	offset := (page - 1) * limit
	var totalCount int64

	var questions []*entity.SurveyQuestion

	if err := r.DB.Model(&entity.SurveyQuestion{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Model(&entity.SurveyQuestion{}).Where("survey_id = ?", surveyId).Limit(limit).Offset(offset).Find(&questions).Error; err != nil {
		return nil, 0, err
	}

	return questions, totalCount, nil
}

// Update implements survey_question_repository.SurveyQuestionRepository.
func (r *SurveyQuestionRepositoryStruct) Update(question *entity.SurveyQuestion) (*entity.SurveyQuestion, error) {
	if err := r.DB.Model(&entity.SurveyQuestion{}).Where("id = ?", question.ID).Updates(question).Error; err != nil {
		return nil, err
	}
	return question, nil
}

// Delete implements survey_question_repository.SurveyQuestionRepository.
func (r *SurveyQuestionRepositoryStruct) Delete(id uint) error {
	if err := r.DB.Where("id = ?", id).Delete(&entity.SurveyQuestion{}).Error; err != nil {
		return err
	}
	return nil
}

// NewSurveyQuestionRepository creates a new instance of SurveyQuestionRepository.
func NewSurveyQuestionRepository(db *gorm.DB) survey_question_repository.SurveyQuestionRepository {
	return &SurveyQuestionRepositoryStruct{db}
}
