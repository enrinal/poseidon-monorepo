package entities

type User struct {
	Model
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
