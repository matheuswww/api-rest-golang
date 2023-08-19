package repository

import (
	"errors"
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"github.com/virussv/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (
	model.UserDomainInterface,*rest_err.RestErr) {
	logger.Info("Init createUser repository",zap.String("journey","createUser"))
	db := ur.databaseConnection
	searchRes,searchErr := ur.FindUser("email",userDomain.GetEmail())
	if searchRes != nil { 
		logger.Error("Error trying to create user",errors.New("duplicated email"),zap.String("journey","createUser"))
		return nil,rest_err.NewConflictError("email already in use")
	}
	if searchErr != nil && searchRes != nil {
		logger.Error("Error trying to create user",searchErr,zap.String("journey","createUser"))
		return nil,searchErr
	}
	query := "INSERT INTO users (email,password,name,age) VALUES (?, ?, ?, ?)"
	result,err := db.Exec(query,userDomain.GetEmail(),userDomain.GetPassword(),userDomain.GetName(),userDomain.GetAge())
	if err != nil {
		logger.Error("Error trying to create user",err,zap.String("journey","createUser"))
		return nil, rest_err.NewInternalServerError("database error")
	}
	id,err := result.LastInsertId()
	if err != nil {
		logger.Error("Error trying to getting id",err,zap.String("journey","createUser"))
		return nil, rest_err.NewInternalServerError("database error")
	}
	userDomain.SetId(uint(id))
	logger.Info("USER INSERTED IN MYSQL DATABASE")
	return userDomain,nil
}