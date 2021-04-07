package db

import "gorm.io/gorm"

func ProvideDefaultDb() *gorm.DB {
	return dbService.DefaultDb
}
