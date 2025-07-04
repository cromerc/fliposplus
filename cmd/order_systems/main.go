package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flipos/internal/adapter/emu_dir"
	"flipos/internal/app"
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/google/go-cmp/cmp"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"io"
	"os"
	"strings"
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

	dirs, err := os.ReadDir(config.RootPath + "/Emu")
	if err != nil {
		log.Fatal().Err(err).Msg("Emu directory does not exist")
	}

	orderF, err := os.OpenFile(config.RootPath+"/order.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		log.Fatal().Err(err).Str("file", "order.txt").Msg("Could not open file")
	}
	defer orderF.Close()

	emuDirs := make(map[string]struct{})
	orderDirs := make(map[string]struct{})
	systemOrder := make([]string, 0)

	buf := bufio.NewReader(orderF)
	for {
		system, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal().Err(err).Msg("Could not read system in order.txt")
		}

		system = strings.TrimSpace(system)

		orderDirs[system] = struct{}{}
		systemOrder = append(systemOrder, system)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			_, err = os.Stat(config.RootPath + "/Emu/" + dir.Name() + "/config.json")
			if err != nil {
				continue
			}

			emuDirs[dir.Name()] = struct{}{}
			log.Trace().Str("dir", dir.Name()).Msg("Found directory")

			if _, exists := orderDirs[dir.Name()]; !exists {
				_, err = orderF.WriteString(dir.Name() + "\n")
				if err != nil {
					log.Fatal().Err(err).Str("file", dir.Name()).Msg("Could not write file")
				}
				log.Trace().Str("dir", dir.Name()).Msgf("Added missing directory")
				orderDirs[dir.Name()] = struct{}{}
			}
		}
	}

	// TODO: decide how to handle inconsistencies between the Emu directories and what is in the order.txt file
	d := cmp.Diff(orderDirs, emuDirs)
	fmt.Println(d)

	i := 0
	for _, system := range systemOrder {
		if _, ok := emuDirs[system]; !ok {
			// The emu directory doesn't exist, skip to the next one
			continue
		}

		systemConf, err := readConfig(config, system)
		if err != nil {
			log.Fatal().Err(err).Msg("Could not read json config")
		}

		addSpaces(systemConf, i, len(systemOrder))

		err = writeConfig(config, system, systemConf)
		if err != nil {
			log.Fatal().Err(err).Msg("Could not write json config")
		}
		log.Info().Str("system", system).Msg("Modified json")
		i++
	}
}

func readConfig(config app.Config, system string) (*emu_dir.Config, error) {
	path := fmt.Sprintf("%s/Emu/%s/config.json", config.RootPath, system)
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg emu_dir.Config

	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	cfg.Label = strings.TrimSpace(cfg.Label)

	return &cfg, nil
}

func writeConfig(config app.Config, system string, cfg *emu_dir.Config) error {
	path := fmt.Sprintf("%s/Emu/%s/config.json", config.RootPath, system)
	b, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, b, 0640)
	if err != nil {
		return err
	}

	return nil
}

// addSpaces adds spaces around the label of the system, this causes the alphabetical order system on the Miyoo Flip to
// move the ones with spaces to the start
func addSpaces(cfg *emu_dir.Config, order, count int) {
	diff := count - order - 1

	var sb strings.Builder

	for i := 0; i < diff; i++ {
		sb.WriteString(" ")
	}

	sb.WriteString(cfg.Label)

	for i := 0; i < diff; i++ {
		sb.WriteString(" ")
	}

	cfg.Label = sb.String()
	cfg.ChineseLabel = sb.String()
}
