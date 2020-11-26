package wallpapers

import (
	"encoding/json"

	"github.com/Vishal-Gupta19/go_gin_website/entity/database"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Wallpaper model
type Wallpaper struct {
	gorm.Model
	Title       string  `gorm:"column:title"`
	Description string  `gorm:"column:description"`
	Image       *string `gorm:"column:image"`
}

// Add all db related operations for this model below

// GetWallpapers : Handler to GET all wallpapers from DB
func GetWallpapers(g *gin.Context) {
	db := database.GetDB()
	defer db.Close()

	var wallpapers []Wallpaper
	db.Find(&wallpapers)
	json.NewEncoder(g.Writer).Encode(wallpapers)
}
