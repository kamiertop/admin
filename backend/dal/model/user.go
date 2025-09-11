package model

type User struct {
	ID       uint32 `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Password string `json:"password,omitempty" db:"password"`
	Email    string `json:"email" db:"email"`
	Gender   string `json:"gender" db:"gender"`
}
