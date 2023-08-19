package service

import (
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"github.com/virussv/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface,string,*rest_err.RestErr ) {
	logger.Info("Init LoginUserServices model",zap.String("journey","LoginUserServices"))

	userDomain.EncryptPassword()

	user,err := ud.findUserByEmailAndPassword(userDomain.GetEmail(),userDomain.GetPassword())
	if err != nil {
		logger.Error("Init LoginUserServices model",err,zap.String("jouney","LoginUserServices"))
		return nil,"",err
	}

	token,err := user.GenerateToken()
	if err != nil {
		return nil,"",err
	}

	return user,token,nil
}