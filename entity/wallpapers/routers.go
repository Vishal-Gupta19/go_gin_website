package wallpapers

/*
 * Collection of all GET/SET apis and handlers for Wallpaper model
 */

import (
	"github.com/gin-gonic/gin"
)

// Mapping paths to api groups

// WallpaperGroup : Modification on Personal wallpapers - apis
func WallpaperGroup(router *gin.RouterGroup) {
	// router.POST("/", AddWallpaper)
	// router.PUT("/login", UpdateWallpaper)
	// router.GET("/", GetWallpaper)
	// router.DELETE("/", DeleteWallpaper)
}

// WallpaperOtherGroup : Save other users wallpapers apis
func WallpaperOtherGroup(router *gin.RouterGroup) {
	// router.POST("/:<something>", addToCollection)
	// router.POST("/:<something>", likeImage)
	// router.POST("/:<something>", downloadImage)
	// router.GET("/:<something>", viewImage)
	// router.DELETE("/:<something>", removeFromCollection)
}
