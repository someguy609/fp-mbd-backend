package entity

import "time"

type Milestone struct {
	MilestoneID       uint      `gorm:"primaryKey;autoIncrement" json:"milestone_id"`
	Title             string    `gorm:"type:varchar(100);not null" json:"title"`
	Description       string    `gorm:"type:text" json:"description"`
	DueDate           time.Time `json:"due_date"`
	Status            string    `gorm:"type:varchar(100);not null" json:"status"`
	CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	ProjectsProjectID uint      `gorm:"not null" json:"projects_project_id"`
	// UsersUserID       string    `gorm:"type:varchar(15);not null" json:"users_user_id"`

	Project Project `gorm:"foreignKey:ProjectsProjectID" json:"project,omitempty"`
	// User    User    `gorm:"foreignKey:UsersUserID" json:"user,omitempty"`
}

func (Milestone) TableName() string {
	return "milestones"
}
