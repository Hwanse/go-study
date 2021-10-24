package user

import "time"

type User struct {
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	CreatedAt 	time.Time	`json:"created_at"`
}

func NewUser(name string, email string) *User {
	return &User{Name: name, Email: email}
}

func (u *User) GenerateCreateTime() {
	u.CreatedAt = time.Now()
}
