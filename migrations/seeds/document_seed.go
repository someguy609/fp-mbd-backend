package seeds

import (
	"encoding/json"
	"fp_mbd/entity"
	"io"
	"os"

	"gorm.io/gorm"
)

func ListDocumentSeeder(db *gorm.DB) error {

	jsonFile, err := os.Open("./migrations/json/documents.json")
	if err != nil {
		return err
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listProject []entity.Project
	if err := json.Unmarshal(jsonData, &listProject); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Project{}); err != nil {
			return err
		}
	}

	for _, data := range listProject {
		if err := db.Create(&data).Error; err != nil {
			return err
		}
	}

	return nil
}
