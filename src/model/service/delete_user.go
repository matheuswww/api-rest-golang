package service

import (
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(id string) (*rest_err.RestErr) {
	logger.Info("Init DeleteUser model",zap.String("journey","DeleteUser"))

	err := ud.userRepository.DeleteUser(id)
	if err != nil {
		logger.Error("Init DeleteUser model",err,zap.String("journey","DeleteUser"))
		return err
	}

	return nil
}