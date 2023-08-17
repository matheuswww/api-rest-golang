package service

import (
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"github.com/virussv/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(queryType string,value string) (model.UserDomainInterface,*rest_err.RestErr) {
	logger.Info("Init find user model",zap.String("journey","FindUser"))
	
	userDomainRepository,err := ud.userRepository.FindUser(queryType,value)
	if err != nil {
		logger.Error("Init find user model",err,zap.String("journey","FindeUsers"))
		return nil,err
	}

	return userDomainRepository,nil
}