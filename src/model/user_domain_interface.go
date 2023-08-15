package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() uint8
	GetName() string
	GetId() uint
	SetId(id uint)

	EncryptPassword()
}

func NewUserDomain(email,password,name string,age uint8,id uint) *userDomain {
	return &userDomain{
		email: email,
		password: password,
		name: name,
		age: age,
		id: id,
	}
}