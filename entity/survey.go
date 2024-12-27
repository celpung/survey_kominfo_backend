package entity

import "time"

type Survey struct {
	ID          uint             `gorm:"primaryKey" json:"id" form:"id"`
	UserID      uint             `json:"user_id" form:"user_id"`
	Author      *User            `gorm:"foreignKey:UserID" json:"author,omitempty"`
	Image       string           `json:"image" form:"image"`
	Title       string           `gorm:"not null" json:"title" form:"title"`
	Slug        string           `gorm:"unique;not null" json:"slug" form:"slug"`
	Key         string           `gorm:"unique;not null" json:"key" form:"key"`
	Status      bool             `json:"status" form:"status"`
	Description string           `json:"description" form:"description"`
	ExpireDate  time.Time        `json:"expire_date" form:"expire_date"`
	Questions   []SurveyQuestion `gorm:"foreignKey:SurveyID" json:"questions"`
	Public      bool             `json:"public"`
	CategoryID  uint             `gorm:"not null" json:"category_id" form:"category_id" binding:"required"`
	Category    *SurveyCategory  `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	CreatedAt   time.Time        `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time       `gorm:"index" json:"deleted_at"`
}

type SurveyCategory struct {
	ID        uint       `gorm:"primaryKey" json:"id" form:"id"`
	Name      string     `gorm:"unique" json:"name" form:"name" binding:"required"`
	Surveys   []Survey   `gorm:"foreignKey:CategoryID" json:"surveys"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

const (
	TypeText     = "text"
	TypeCheckbox = "checkbox"
	TypeRadio    = "radio"
	TypeDropdown = "dropdown"
	TypeLocation = "location"
)

type SurveyQuestion struct {
	ID          uint           `gorm:"primaryKey" json:"id" form:"id"`
	SurveyID    uint           `json:"survey_id" form:"survey_id"`
	Survey      *Survey        `gorm:"foreignKey:SurveyID" json:"survey"`
	Question    string         `gorm:"type:varchar(255);not null" json:"question" form:"question" binding:"required"`
	Type        string         `gorm:"type:varchar(50);not null" json:"type" form:"type" binding:"required"`
	Description string         `gorm:"type:text" json:"description" form:"description"`
	Data        string         `gorm:"type:text" json:"data" form:"data"`
	AllowImage  bool           `json:"allow_image" form:"allow_image"`
	Answers     []SurveyAnswer `gorm:"foreignKey:QuestionID" json:"answers"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time     `gorm:"index" json:"deleted_at"`
}

// kategori survey desnt needed

type SurveyAnswer struct {
	ID         uint            `gorm:"primaryKey" json:"id" form:"id"`
	SurveyID   uint            `json:"survey_id" form:"survey_id"`
	Survey     *Survey         `gorm:"foreignKey:SurveyID" json:"survey"`
	QuestionID uint            `json:"question_id" form:"question_id"`
	Question   *SurveyQuestion `gorm:"foreignKey:QuestionID" json:"question"`
	Answer     string          `gorm:"type:text" json:"answer" form:"answer"`
	Image      string          `json:"image" form:"image"`
	CreatedAt  time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  *time.Time      `gorm:"index" json:"deleted_at"`
}

// UserID     uint           `json:"user_id" form:"user_id"`
// 	Surveyor   User           `gorm:"foreignKey:UserID" json:"author"`

// need table unverified answer
// need table
// error response
