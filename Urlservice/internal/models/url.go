package models

import (
	"url-service/internal/utils"
	"github.com/gofrs/uuid"
)


type URL struct {
	Id       string `json:"id,omitempty"`
	Redirect string `json:"redirect,omitempty"`
	TVF      string `json:"TVF,omitempty"`
	Random   bool   `json:"random,omitempty"`
	UserID   string	`json:"user_id,omitempty"`
}



func(u *URL) PrepareBeforeInsert(){
	if u.TVF == "" {
		u.TVF = "TVF" + utils.GenerateRandomString()
		u.Random = true
	}
	u.TVF = "TVF" + u.TVF
}

func(u *URL) GenID()string{
	id, _ :=  uuid.DefaultGenerator.NewV1()
	u.Id =  id.String()
	return u.Id
}