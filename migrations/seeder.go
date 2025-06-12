package migrations

import (
	"fp_mbd/migrations/seeds"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListUserSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListProjectSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListDocumentSeeder(db); err != nil {
		return err
	}
	if err := seeds.ListMilestoneSeeder(db); err != nil {
		return err
	}
	if err := seeds.ListProjectMemberSeeder(db); err != nil {
		return err
	}

	return nil
}
