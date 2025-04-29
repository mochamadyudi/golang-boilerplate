package config

import "gorm.io/gorm"

func GormConfig() *gorm.Config {
	// config gorm in here
	return &gorm.Config{
		PrepareStmt: true,
	}
}
