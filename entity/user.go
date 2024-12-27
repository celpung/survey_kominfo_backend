package entity

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `json:"name" binding:"required"`
	Username  string     `gorm:"unique" json:"username" binding:"required"`
	Password  string     `gorm:"not null" json:"password" binding:"required"`
	Active    bool       `gorm:"default:0" json:"active"`
	Role      uint       `gorm:"not null;default:2" json:"role"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

// Phone     string     `gorm:"not null" json:"phone"`
// SurveyAnswer []SurveyAnswer `gorm:"foreignKey:UserID" json:"survey_answer"`
// Groups       []Group        `gorm:"many2many:group_users" json:"groups"`
type Role struct {
	// menghubungkan antara user dan survey
	// user ini punya akses apa ke survey mana
	// role ini berguna untuk
	// kategori survey
}

type Group struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"unique" json:"name" binding:"required"`
	Pic       uint       `json:"pic"`
	Users     []User     `gorm:"many2many:group_users" json:"users"`
	Surveys   []Survey   `gorm:"foreignKey:UserID" json:"surveys"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

type UserUpdate struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	Role     uint   `json:"role"`
}

type UserHttpResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Active   bool   `json:"active"`
	Role     uint   `json:"role"`
}
