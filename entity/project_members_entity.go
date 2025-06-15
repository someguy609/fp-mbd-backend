package entity

import "time"

type ProjectMember struct {
	ProjectMemberID uint      `gorm:"primaryKey;autoIncrement" json:"project_member_id"`
	RoleProject     string    `gorm:"column:role_project;type:varchar(10);not null" json:"role_project"` // Menggunakan gorm:"column" untuk nama kolom yang berbeda
	JoinedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"joined_at"`
	UserID          string    `gorm:"type:varchar(15);not null" json:"user_id"`
	ProjectID       uint      `gorm:"not null" json:"project_id"`

	User    User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Project Project `gorm:"foreignKey:ProjectID" json:"project,omitempty"`
}

func (ProjectMember) TableName() string {
	return "project_members"
}
