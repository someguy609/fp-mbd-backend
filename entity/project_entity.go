package entity

import (
	"time"

	"github.com/lib/pq"
)

type ProjectStatus string

const (
	StatusPlanning  ProjectStatus = "PLANNING"
	StatusOngoing   ProjectStatus = "ONGOING"
	StatusCompleted ProjectStatus = "COMPLETED"
	StatusSuspended ProjectStatus = "SUSPENDED"
)

type Project struct {
	ProjectID   uint           `gorm:"primaryKey;autoIncrement" json:"project_id"`
	Title       string         `gorm:"type:varchar(100);not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Status      ProjectStatus  `gorm:"type:varchar(20);not null" json:"status"`
	StartDate   time.Time      `gorm:"type:timestamp,default:CURRENT_TIMESTAMP" json:"start_date"`
	EndDate     time.Time      `gorm:"type:timestamp" json:"end_date"`
	Categories  pq.StringArray `gorm:"type:text[]" json:"categories"`
	CreatedAt   time.Time      `gorm:"type:timestamp,default:CURRENT_TIMESTAMP" json:"created_at"`

	ProjectMembers []ProjectMember `gorm:"foreignKey:ProjectsProjectID" json:"project_members,omitempty"`
	Milestones     []Milestone     `gorm:"foreignKey:ProjectsProjectID" json:"milestones,omitempty"`
	Documents      []Document      `gorm:"foreignKey:ProjectsProjectID" json:"documents,omitempty"`
}

// TableName secara eksplisit memberi tahu GORM nama tabel untuk struct Project
func (Project) TableName() string {
	return "projects"
}
