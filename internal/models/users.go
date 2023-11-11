package models

type User struct {
	UserID   int    `json:"user_id,omitempty"`
	Username string `json:"login,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
