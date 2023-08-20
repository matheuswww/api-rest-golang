package service

import (
	"github.com/matheuswww/api-rest-golang/src/configuration/logger"
	"github.com/matheuswww/api-rest-golang/src/configuration/rest_err"
	"github.com/matheuswww/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string,userDomain model.UserDomainInterface) (*rest_err.RestErr) {
	logger.Info("Init UpdateUser model",zap.String("journey","UpdateUser"))

	err := ud.userRepository.UpdateUser(userId,userDomain)
	if err != nil {
		logger.Error("error trying update user model",err,zap.String("jouney","updateUser"))
		return err
	}

	return  nil
}