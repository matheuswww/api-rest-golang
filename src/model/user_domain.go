package model

import (
	"crypto/sha512"
	"encoding/hex"
)

func NewUserDomain(email,password,name string,age int8,) *userDomain {
	return &userDomain{
		email,password,name,age,
	}
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	EncryptPassword()
}

type userDomain struct {
	email string
	password string
	name string
	age int8
}

func (ud *userDomain) GetEmail() string{
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetName() string{
	return ud.name
}
func (ud *userDomain) GetAge() int8{
	return ud.age
}

func (ud *userDomain) EncryptPassword() {
	hash := sha512.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))	
}