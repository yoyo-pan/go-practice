package models

type User struct {
	Username	string	`binding:"required" json:"username"`
	Email		string	`binding:"required" json:"email"`
}
