package survey_repository_implementation

import (
	survey_repository "github.com/celpung/gocleanarch/domain/survey/repository"
	"github.com/celpung/gocleanarch/entity"
	"gorm.io/gorm"
)

type SurveyRepositoryStruct struct {
	DB *gorm.DB
}

// Create implements survey_repository.SurveyRepositoryInterface.
func (r *SurveyRepositoryStruct) Create(survey *entity.Survey) (*entity.Survey, error) {
	// insert data into databse
	if err := r.DB.Create(survey).Error; err != nil {
		return nil, err
	}

	return survey, nil
}

// Read implements survey_repository.SurveyRepositoryInterface.
func (r *SurveyRepositoryStruct) Read() ([]*entity.Survey, error) {
	// get all data from database
	var surveys []*entity.Survey
	if err := r.DB.Preload("Questions").Find(&surveys).Error; err != nil {
		return nil, err
	}

	return surveys, nil
}

// ReadByID implements survey_repository.SurveyRepositoryInterface.
func (r *SurveyRepositoryStruct) ReadByID(id uint) (*entity.Survey, error) {
	// read survey by id
	var survey entity.Survey
	if err := r.DB.Preload("Questions").First(&survey, id).Error; err != nil {
		return nil, err
	}

	return &survey, nil
}

// Update implements survey_repository.SurveyRepositoryInterface.
func (r *SurveyRepositoryStruct) Update(survey *entity.Survey) (*entity.Survey, error) {
	// update the survey by id
	if err := r.DB.Model(&entity.Survey{}).Where("id = ?", survey.ID).Updates(survey).Error; err != nil {
		return nil, err
	}

	return survey, nil
}

// Delete implements survey_repository.SurveyRepositoryInterface.
func (r *SurveyRepositoryStruct) Delete(id uint) error {
	// delete by id
	if err := r.DB.Delete(&entity.Survey{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewSurveyRepositry(db *gorm.DB) survey_repository.SurveyRepositoryInterface {
	return &SurveyRepositoryStruct{DB: db}
}
