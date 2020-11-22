package model

// type tag int

// const (
// 	Nature tag = iota
// 	People
// 	Architecture
// 	Experimental
// 	Fashion
// 	Film
// 	Health
// 	Interiors
// 	Technology
// 	Travel
// 	Animals
// 	Art
// 	Culture
// 	Athletics
// 	History
// )

// Wallpaper : Model for Wallpaper
// It is inheriting a User model
type Wallpaper struct {
	Title       string `json:"title" binding:"min=2,max=50"`
	Description string `json:"description" binding:"max=200"`
	URL         string `json:"url" binding:"required,url"`
	//	Tag         tag    `json:"tag" binding:"max=200"`
	Author User `json:"author" binding:"required"`
}

// User : Model for user credentials
type User struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `form:"password" json:"password" binding:"required"`
}
