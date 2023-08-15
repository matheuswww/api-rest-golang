package model

type userDomain struct {
	email string
	password string
	name string
	age uint8
	id uint
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
func (ud *userDomain) GetAge() uint8{
	return ud.age
}
func (ud *userDomain) GetId() uint {
	return ud.id
}
func (ud *userDomain) SetId(id uint) {
	ud.id = id 
}