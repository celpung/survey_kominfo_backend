package entity

import "time"

type Survey struct {
	ID          uint             `gorm:"primaryKey" json:"id" form:"id"`
	UserID      uint             `json:"user_id" form:"user_id"`
	Author      User             `gorm:"foreignKey:UserID" json:"author"`
	Image       string           `json:"image" form:"image" binding:"required"`
	Title       string           `gorm:"unique" json:"title" form:"title" binding:"required"`
	Slug        string           `gorm:"unique" json:"slug" form:"slug" binding:"required"`
	Status      bool             `json:"status" form:"status" binding:"required"`
	Description string           `json:"description" form:"description"`
	ExpireDate  time.Time        `json:"expire_date" form:"expire_date" binding:"required"`
	Questions   []SurveyQuestion `gorm:"foreignKey:SurveyID" json:"questions"`
	CreatedAt   time.Time        `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time       `gorm:"index" json:"deleted_at"`
}

const (
	TypeText     = "text"
	TypeCheckbox = "checkbox"
	TypeRadio    = "radio"
	TypeDropdown = "dropdown"
)

type SurveyQuestion struct {
	ID          uint       `gorm:"primaryKey" json:"id" form:"id"`
	SurveyID    uint       `json:"survey_id" form:"survey_id"`
	Survey      Survey     `gorm:"foreignKey:SurveyID" json:"survey"`
	Question    string     `gorm:"unique" json:"question" form:"question" binding:"required"`
	Type        string     `json:"type" form:"type" binding:"required"`
	Description string     `json:"description" form:"description"`
	Data        string     `json:"data" form:"data"`
	AllowImage  bool       `json:"allow_image" form:"allow_image"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
}

type QuestionOptions struct {
	Options []string `json:"options"`
}

type SurveyAnswer struct {
	ID         uint           `gorm:"primaryKey" json:"id" form:"id"`
	SurveyID   uint           `json:"survey_id" form:"survey_id"`
	Survey     Survey         `gorm:"foreignKey:SurveyID" json:"survey"`
	QuestionID uint           `json:"question_id" form:"question_id"`
	Question   SurveyQuestion `gorm:"foreignKey:QuestionID" json:"question"`
	Answer     string         `json:"answer" form:"answer"`
	Image      string         `json:"image" form:"image"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  *time.Time     `gorm:"index" json:"deleted_at"`
}
