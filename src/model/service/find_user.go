package service

import (
	"github.com/matheuswww/api-rest-golang/src/configuration/logger"
	"github.com/matheuswww/api-rest-golang/src/configuration/rest_err"
	"github.com/matheuswww/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(queryType string,value string) (model.UserDomainInterface,*rest_err.RestErr) {
	logger.Info("Init find user model",zap.String("journey","FindUser"))
	userDomainRepository,err := ud.userRepository.FindUser(queryType,value)
	if err != nil {
		logger.Error("error trying find user model",err,zap.String("jouney","findUser"))
		return nil,err
	}
	return userDomainRepository,nil
}

func (ud *userDomainService) findUserByEmailAndPassword(email string,password string) (model.UserDomainInterface,*rest_err.RestErr) {
	user,err := ud.userRepository.FindUserByEmailAndPassword(email,password)
	if err != nil {
		logger.Error("error trying find user model",err,zap.String("jouney","findUserByEmailAndPassword"))
		return nil,err
	}
	return user,nil
}