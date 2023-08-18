package service

import (
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"github.com/virussv/api-rest-golang/src/model"
	"github.com/virussv/api-rest-golang/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository,) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface,*rest_err.RestErr)
	UpdateUser(string,model.UserDomainInterface) *rest_err.RestErr
	FindUser(string,string) (model.UserDomainInterface,*rest_err.RestErr)
	LoginUserServices(model.UserDomainInterface) (model.UserDomainInterface,*rest_err.RestErr)
	DeleteUser(string) (*rest_err.RestErr)
}