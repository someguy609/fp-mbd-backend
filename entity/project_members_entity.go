package entity

import "time"

type ProjectMember struct {
	ProjectMemberID   uint      `gorm:"primaryKey;autoIncrement" json:"project_member_id"`
	RoleProject       string    `gorm:"column:role_project;type:varchar(10);not null" json:"role_project"` // Menggunakan gorm:"column" untuk nama kolom yang berbeda
	IsActive          bool      `gorm:"default:false" json:"is_active"`                                    // Menambahkan is_active untuk status keanggotaan
	JoinedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"joined_at"`
	ProjectsProjectID uint      `gorm:"not null" json:"projects_project_id"`
	UsersUserID       string    `gorm:"type:varchar(15);not null" json:"users_user_id"`

	Project Project `gorm:"foreignKey:ProjectsProjectID" json:"project,omitempty"`
	User    User    `gorm:"foreignKey:UsersUserID" json:"user,omitempty"`
}

func (ProjectMember) TableName() string {
	return "project_members"
}
