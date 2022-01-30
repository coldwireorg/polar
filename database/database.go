package database

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Export conn
var DB *gorm.DB

// Run database
func Connect(path string) {
	conn, err := gorm.Open(sqlite.Open(path+"/polar.db?mode=memory"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Info().Msg("Connected to database!")

	DB = conn
}
