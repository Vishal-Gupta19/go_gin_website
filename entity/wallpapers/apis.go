package wallpapers

/*
 * Collection of all GET/SET handlers for Wallpaper model
 */

import (
	"encoding/json"

	"github.com/Vishal-Gupta19/go_gin_website/entity"
	"github.com/gin-gonic/gin"
)

// GetWallpapers : Handler to GET all wallpapers from DB
func GetWallpapers(g *gin.Context) {
	db := entity.GetDB()
	defer db.Close()

	var wallpapers []Wallpaper
	db.Find(&wallpapers)
	json.NewEncoder(g.Writer).Encode(wallpapers)
}
