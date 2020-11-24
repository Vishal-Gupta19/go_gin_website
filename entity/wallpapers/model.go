package wallpapers

import (
	"github.com/jinzhu/gorm"
)

// Wallpaper model
type Wallpaper struct {
	gorm.Model
	Title       string  `gorm:"column:title"`
	Description string  `gorm:"column:description"`
	Image       *string `gorm:"column:image"`
}
