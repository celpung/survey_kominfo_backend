package survey_usecase_implementation

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	survey_repository "github.com/celpung/gocleanarch/domain/survey/repository"
	survey_usecase "github.com/celpung/gocleanarch/domain/survey/usecase"
	"github.com/celpung/gocleanarch/entity"
)

type SurveyUsecaseStruct struct {
	repository survey_repository.SurveyRepositoryInterface
}

// Create implements survey_usecase.SurveyUsecaseInterface.
func (u *SurveyUsecaseStruct) Create(survey *entity.Survey) (*entity.Survey, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var key strings.Builder
	for i := 0; i < 10; i++ {
		randomIndex := rng.Intn(len(charset))
		key.WriteByte(charset[randomIndex])
	}

	survey.Key = key.String()

	// Regular expression to allow only letters, numbers, and hyphens (-)
	re := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	if !re.MatchString(survey.Slug) {
		return nil, fmt.Errorf("slug hanya boleh huruf, angka, dan strip (-)")
	}

	return u.repository.Create(survey)
}

// Read implements survey_usecase.SurveyUsecaseInterface.
func (u *SurveyUsecaseStruct) Read(page, limit int) ([]*entity.Survey, int64, error) {
	return u.repository.Read(page, limit)
}

func (u *SurveyUsecaseStruct) ReadByID(id uint) (*entity.Survey, error) {
	return u.repository.ReadByID(id)
}

func (u *SurveyUsecaseStruct) ReadBySlug(slug string) (*entity.Survey, error) {
	return u.repository.ReadBySlug(slug)
}

// Update implements survey_usecase.SurveyUsecaseInterface.
func (u *SurveyUsecaseStruct) Update(survey *entity.Survey) (*entity.Survey, error) {
	// update survey
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var key strings.Builder
	for i := 0; i < 10; i++ {
		randomIndex := rng.Intn(len(charset))
		key.WriteByte(charset[randomIndex])
	}

	survey.Key = key.String()

	survey.Key = key.String()
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
