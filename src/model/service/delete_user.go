package service

import (
	"github.com/matheuswww/api-rest-golang/src/configuration/logger"
	"github.com/matheuswww/api-rest-golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(id string) (*rest_err.RestErr) {
	logger.Info("Init DeleteUser model",zap.String("journey","DeleteUser"))

	err := ud.userRepository.DeleteUser(id)
	if err != nil {
		logger.Error("error trying delete user model",err,zap.String("jouney","deleteUser"))
		return err
	}

	return nil
}