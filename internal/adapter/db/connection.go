package db

import (
	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open(log *zerolog.Logger, file string) (*gorm.DB, error) {
	log.Trace().Str("file", file).Msg("Opening database")
	gdb, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gdb, nil
}
