package app

type Config struct {
	RootPath string `env:"ROOT_PATH" envDefault:"/mnt/SDCARD"`
}
