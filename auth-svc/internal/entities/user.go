package entities

type User struct {
	Model
	Name     string `json:"name" gorm:"unique"`
	Phone    string `json:"phone" gorm:"unique"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
