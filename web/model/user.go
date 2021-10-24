package model

import "time"

type User struct {
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	CreatedAt 	time.Time	`json:"created_at"`
}

func (u User) NewUser(name, email string) User {
	u.Name = name
	u.Email = email
	return u
}

func (u *User) GenerateCreateTime() {
	u.CreatedAt = time.Now()
}
