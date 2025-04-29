package main

import (
	"errors"
	"flipos/internal/adapter/db"
	"flipos/internal/app"
	"flipos/internal/domain"
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"os"
	"regexp"
)

func main() {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	var config app.Config
	config, err = env.ParseAs[app.Config]()
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(config.RootPath+"/log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	log := zerolog.New(f).With().Timestamp().Logger()

	dirs, err := os.ReadDir(config.RootPath + "/Roms")
	if err != nil {
		log.Fatal().Err(err).Msg("Roms directory does not exist")
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			log.Trace().Str("dir", dir.Name()).Msg("Cleaning directory")
			rd := domain.NewRomDatabase(&config, dir.Name())

			_, err = os.Stat(rd.Path())
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					continue
				}
				log.Error().Err(err).Msg("Could not stat Rom database")
			}

			err = cleanNames(&log, *rd)
			if err != nil {
				log.Error().Str("dir", rd.CodeName).Err(err).Msg("Could not clean names")
			}
		}
	}
}

func cleanNames(log *zerolog.Logger, database domain.RomDatabase) error {
	gdb, err := db.Open(log, database.Path())
	if err != nil {
		return fmt.Errorf("could not open database %s: %w", database.Path(), err)
	}

	log.Info().Str("db", database.CodeName).Msg("Finding roms")
	var roms []db.Rom
	result := gdb.Table(database.TableName).Find(&roms)
	if result.Error != nil {
		return err
	}

	for _, rom := range roms {
		log.Info().Str("rom", rom.Disp).Msg("Cleaning up rom name")

		if rom.Type != 0 {
			// Not a rom file, skip
			continue
		}

		removeParenthesis(&rom)
		moveArticle(&rom)

		log.Info().Str("rom", rom.Disp).Msg("Name cleaned")
		gdb.Table(database.TableName).Save(rom)
	}

	return nil
}

func moveArticle(rom *db.Rom) {
	r := regexp.MustCompile(`,\s(The|A|An)`)
	article := r.FindStringSubmatch(rom.Disp)
	if len(article) > 1 {
		newRomName := r.ReplaceAllString(rom.Disp, "")
		rom.Disp = article[1] + " " + newRomName
		rom.Pinyin = article[1] + " " + newRomName
		rom.Cpinyin = article[1] + " " + newRomName
	}
}

func removeParenthesis(rom *db.Rom) {
	r := regexp.MustCompile(`\s*\(.*\)\s*`)
	newRomName := r.ReplaceAllString(rom.Disp, "")
	rom.Disp = newRomName
	rom.Pinyin = newRomName
	rom.Cpinyin = newRomName
}
