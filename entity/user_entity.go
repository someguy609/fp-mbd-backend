package entity

import (
	"fp_mbd/helpers"

	"gorm.io/gorm"
)

type User struct {
	UserID      string `gorm:"type:varchar(15);primaryKey" json:"user_id"`
	Name        string `gorm:"type:varchar(50);not null" json:"name"`
	Email       string `gorm:"type:varchar(50);not null;unique" json:"email"`
	Role        string `gorm:"type:varchar(10);not null" json:"role"`
	ContactInfo string `gorm:"type:varchar(100)" json:"contact_info"`
	Password    string `gorm:"type:varchar(255);not null" json:"-"`

	ProjectMembers []ProjectMember `gorm:"foreignKey:UserID" json:"project_members,omitempty"`
	Documents      []Document      `gorm:"foreignKey:UsersUserID" json:"documents,omitempty"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	if u.Password != "" {
		u.Password, err = helpers.HashPassword(u.Password)
		if err != nil {
			return err
		}
	}

	if u.Role == "" {
		u.Role = "mahasiswa"
	}

	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	if u.Password != "" {
		u.Password, err = helpers.HashPassword(u.Password)
		if err != nil {
			return err
		}
	}
	return nil
}
