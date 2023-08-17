package service

import (
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"github.com/virussv/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string,userDomain model.UserDomainInterface) (*rest_err.RestErr) {
	logger.Info("Init UpdateUser model",zap.String("journey","UpdateUser"))

	err := ud.userRepository.UpdateUser(userId,userDomain)
	if err != nil {
		logger.Error("Init UpdateUser model",err,zap.String("jouney","UpdateUser"))
		return err
	}

	return  nil
}