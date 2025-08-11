package user

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) GetLastName() string{
	return u.lastName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) OutputUserInfo(){
	fmt.Println(u.firstName, u.lastName)
}

type Admin struct{
	email string
	password string
	User
}

func NewAdmin(email, password string) (*Admin, error){
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == ""{
		return nil, errors.New("email and password cannot be empty")
	}
	return &Admin{
		email: email,
		password: password,
	}, nil
}

func New(firstName, lastName, birthDate string)(*User, error){
	if strings.TrimSpace(firstName) == "" || strings.TrimSpace(lastName) == "" || strings.TrimSpace(birthDate) == ""{
		return nil, errors.New("first name, last name, birth date cannot be empty")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}