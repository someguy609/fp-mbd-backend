package migrations

import (
	"os"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	sqlScript, err := os.ReadFile("database.sql")

	if err != nil {
		return err
	}
	if err := db.Exec(string(sqlScript)).Error; err != nil {
		return err
	}

	return nil
}
