package db

type Rom struct {
	ID      int `gorm:"primaryKey"`
	Disp    string
	Path    string
	Imgpath string
	Type    int
	Ppath   string
	Pinyin  string
	Cpinyin string
}
