package repository

import "gorm.io/gorm"

func EnsureUUIDExtension(db *gorm.DB) error {
	var result string
	err := db.Raw("SELECT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'uuid-ossp')").Row().Scan(&result)
	if err != nil {
		return err
	}

	if result != "t" {
		if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
			return err
		}
	}
	return nil
}
