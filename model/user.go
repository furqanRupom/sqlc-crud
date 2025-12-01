package model

type User struct {
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}
type RUser struct {
	ID    int `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
