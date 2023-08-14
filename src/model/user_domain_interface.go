package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	EncryptPassword()
}

func NewUserDomain(email,password,name string,age int8) *userDomain {
	return &userDomain{
		email,password,name,age,
	}
}