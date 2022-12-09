package model

import (
	"strings"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
}

// Sanitize password
func (u *User) SanitizePassword() {
	u.Password = ""
}
func(u *User) GenID()string{
	id, _ :=  uuid.DefaultGenerator.NewV1()
	u.ID =  id.String()
	return u.ID
}
func(u *User) GetFirstName() string{
	return u.FirstName 
}
func(u *User) GetLastName() string{
	return u.LastName 
}
func(u *User) GetEmail() string{
	return u.Email
}
func(u *User) GetRole() string{
	return u.Role
}
func(u *User) GetPassword() string{
	return u.Password
}

// Hash user password with bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Compare user password and payload
func (u *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// Prepare user for register
func (u *User) PrepareCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}

	if u.Role != "" {
		u.Role = strings.ToLower(strings.TrimSpace(u.Role))
	}

	return nil
}
