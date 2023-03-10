package pkg

type ImageConfig struct {
	URL    string `json:"url" gorm:"column:url"`
	Width  int    `json:"width" gorm:"column:width"`
	Height int    `json:"height" gorm:"column:height"`
}

type OffSet struct {
	X int `json:"x" gorm:"column:offset_x"`
	Y int `json:"y" gorm:"column:offset_y"`
}

type OffSetConfig struct {
	OffSet `gorm:"embedded"`
	Prefix string `json:"prefix" gorm:"column:prefix"`
	Suffix string `json:"suffix" gorm:"column:suffix"`
}

type FontConfig struct {
	Family string `json:"family" gorm:"column:family"`
	Size   int    `json:"size" gorm:"column:size"`
	Color  string `json:"color" gorm:"column:color"`
	Layout string `json:"layout" gorm:"column:layout"`
	Style  string `json:"style" gorm:"column:style"`
}

type ImageRenderConfig struct {
	ID          uint `gorm:"primaryKey"`
	ImageConfig `json:"image" gorm:"embedded;embeddedPrefix:image_"`
	TextOffSet  OffSet       `json:"textOffSet" gorm:"embedded;embeddedPrefix:text_"`
	Row         OffSetConfig `json:"row" gorm:"embedded;embeddedPrefix:row_"`
	Col         OffSetConfig `json:"col" gorm:"embedded;embeddedPrefix:row_"`
	Font        FontConfig   `json:"font" gorm:"embedded;embeddedPrefix:font_"`
}

func (ImageRenderConfig) TableName() string {
	return "image_render_config"
}
