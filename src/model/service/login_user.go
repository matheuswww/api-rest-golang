package service

import (
	"github.com/matheuswww/api-rest-golang/src/configuration/logger"
	"github.com/matheuswww/api-rest-golang/src/configuration/rest_err"
	"github.com/matheuswww/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface,string,*rest_err.RestErr ) {
	logger.Info("Init LoginUserServices model",zap.String("journey","LoginUserServices"))

	userDomain.EncryptPassword()

	user,err := ud.findUserByEmailAndPassword(userDomain.GetEmail(),userDomain.GetPassword())
	if err != nil {
		logger.Error("error trying login user model",err,zap.String("jouney","loginUserServices"))
		return nil,"",err
	}

	token,err := user.GenerateToken()
	if err != nil {
		return nil,"",err
	}

	return user,token,nil
}