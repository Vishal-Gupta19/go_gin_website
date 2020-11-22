package wallpapers


// Wallpaper model
type Wallpaper struct {
	ID           uint    `gorm:"primary_key"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Image        *string `gorm:"column:image"`
}