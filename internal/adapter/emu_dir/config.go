package emu_dir

type Config struct {
	Label         string       `json:"label,omitempty"`
	ChineseLabel  string       `json:"label.ch.lang,omitempty"`
	Icon          string       `json:"icon,omitempty"`
	SelectIcon    string       `json:"iconsel,omitempty"`
	SmallIcon     string       `json:"iconsmall,omitempty"`
	ListIcon      string       `json:"iconlist,omitempty"`
	Background    string       `json:"background,omitempty"`
	EffectScript  string       `json:"effectsh,omitempty"`
	ThemeColor    string       `json:"themecolor,omitempty"`
	Launch        string       `json:"launch,omitempty"`
	RomPath       string       `json:"rompath,omitempty"`
	ImagePath     string       `json:"imgpath,omitempty"`
	GameList      string       `json:"gamelist,omitempty"`
	UseSwap       int          `json:"useswap,omitempty"`
	ShortName     int          `json:"shortname,omitempty"`
	HideBios      int          `json:"hidebios,omitempty"`
	ExtensionList string       `json:"extensionlist,omitempty"`
	LaunchList    []LaunchList `json:"launchlist,omitempty"`
}

type LaunchList struct {
	Name   string `json:"name,omitempty"`
	Launch string `json:"launch,omitempty"`
}
