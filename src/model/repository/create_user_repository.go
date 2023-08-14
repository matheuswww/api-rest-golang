package repository

import (
	"log"

	"github.com/virussv/api-rest-golang/src/configuration/database/mysql"
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"github.com/virussv/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (
	model.UserDomainInterface,*rest_err.RestErr) {
	logger.Info("Init createUser repository",zap.String("journey","createUser"))

	db,err := mysql.NewMysqlConnection()
	if err != nil {
		log.Fatal("ERROR TO CONNECT MYSQL DB")
	}
	defer db.Close()
	query := "INSERT INTO users (email,name,password,age) VALUES (?, ?, ?, ?)"
    _, err = db.Exec(query,userDomain.GetEmail(),userDomain.GetName(),userDomain.GetPassword(),userDomain.GetAge())
    if err != nil {
			logger.Error("Error trying to create user",err,zap.String("journey","createUser"))
    	return nil, rest_err.NewInternalServerError("database error")
    }
	logger.Info("USER INSERTED IN MYSQL DATABASE")
	return userDomain,nil
}