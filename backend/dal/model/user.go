package model

type User struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}
