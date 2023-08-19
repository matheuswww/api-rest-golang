package service

import (
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"github.com/virussv/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface,*rest_err.RestErr ) {
	logger.Info("Init createUser model",zap.String("journey","createUser"))

	userDomain.EncryptPassword()

	userDomainRepository,err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("error trying create user model",err,zap.String("jouney","createUser"))
		return nil,err
	}

	return userDomainRepository,nil
}