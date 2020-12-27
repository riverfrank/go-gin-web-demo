package dto

type UserCreateDTO struct {
	//jsonName
	Email    string `gorm:"column:email;" json:"email"`
	Name     string `gorm:"column:name;" json:"name"`
	Password string `gorm:"column:password;" json:"password"`
}
