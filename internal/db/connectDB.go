package db

import (
	//"fmt"
	"os"
	"tzAvito/internal/model"
	"tzAvito/internal/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}

	if err := repository.EnsureUUIDExtension(DB); err != nil {
		return
	}

	DB.Exec(`DO $$ BEGIN
                IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'organization_type') THEN
                    CREATE TYPE organization_type AS ENUM ('IE', 'LLC', 'JSC');
                END IF;
             END $$;`)

	err = DB.AutoMigrate(&model.User{}, &model.Organization{}, &model.Tender{}, &model.Bid{}, &model.TenderVersion{})

	if err != nil {
		panic("Failed to migrate DB schemas")
	}

}

func insertInitialData() error { //мне пришлось вот так захардкодить, вроде бы отдельно для этого эндпоинт на сваггере не просили
	var userCount int64
	if err := DB.Model(&model.User{}).Count(&userCount).Error; err != nil {
		return err
	}

	if userCount == 0 {
		users := []model.User{
			{ID: "550e8400-e29b-41d4-a716-446655440000", Username: "User"},
		}
		if err := DB.Create(&users).Error; err != nil {
			return err
		}
	}

	return nil
}
