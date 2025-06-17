package migrations

import (
	"fp_mbd/migrations/seeds"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	// Import the seeders from the seeds package
	seeders := []func(*gorm.DB) error{
		seeds.ListUserSeeder,
		seeds.ListProjectSeeder,
		seeds.ListDocumentSeeder,
		seeds.ListMilestoneSeeder,
		seeds.ListProjectMemberSeeder,
	}

	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			return err
		}
	}

	return nil
}
