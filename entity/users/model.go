package users

// User : Model for user credentials
type User struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	ProfilePic        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
	FirstName 	 string  `gorm:"column:firstname"`
	LastName     string  `gorm:"column:lastname"`	
	City         string  `gorm:"column:city"`

}
