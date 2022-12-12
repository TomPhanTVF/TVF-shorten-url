package models


type User struct {
	Id 			string 	`json:"id"`
	UserName 	string  `json:"user_name"`
	UserEmail 	string  `json:"user_email"`
}