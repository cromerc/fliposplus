package domain

import "flipos/internal/app"

type RomDatabase struct {
	config    *app.Config
	CodeName  string
	TableName string
	FileName  string
}

func (rd *RomDatabase) Path() string {
	return rd.config.RootPath + "/Roms/" + rd.CodeName + "/" + rd.FileName
}

func NewRomDatabase(config *app.Config, codename string) *RomDatabase {
	return &RomDatabase{
		config:    config,
		CodeName:  codename,
		TableName: codename + "_roms",
		FileName:  codename + "_cache6.db",
	}
}
