package entity

import "time"

type Document struct {
	DocumentID        uint      `gorm:"primaryKey;autoIncrement" json:"document_id"`
	Title             string    `gorm:"type:varchar(100);not null" json:"title"`
	FileURL           string    `gorm:"type:varchar(255);not null" json:"file_url"`
	DocumentType      string    `gorm:"type:varchar(100)" json:"document_type"`
	CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	ProjectsProjectID uint      `gorm:"not null" json:"projects_project_id"`
	UsersUserID       string    `gorm:"type:varchar(15);not null" json:"users_user_id"`

	Project Project `gorm:"foreignKey:ProjectsProjectID" json:"project,omitempty"`
	User    User    `gorm:"foreignKey:UsersUserID" json:"user,omitempty"`
}

func (Document) TableName() string {
	return "documents"
}
