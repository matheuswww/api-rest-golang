package repository

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/virussv/api-rest-golang/src/configuration/database/mysql"
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"github.com/virussv/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (us *userRepository) FindUser(queryType string,value string) (model.UserDomainInterface,*rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository",zap.String("journey","FindeUserByEmail"))

	db,err := mysql.NewMysqlConnection()
	if err != nil {
		return nil, rest_err.NewInternalServerError("database error")
	}
	defer db.Close()
	var row *sql.Row
	switch queryType  {
		case "id":
			query := "SELECT email,password,name,age,id FROM users WHERE id = ?"
			res,err := strconv.ParseUint(value,10,64)
			if err != nil {
				logger.Error("Error trying convert int",err,zap.String("journey","FindUser"))
				return nil,rest_err.NewBadRequestError("invalid id")
			}
			row = db.QueryRow(query,res)
		case "email":
			query := "SELECT email,password,name,age,id FROM users WHERE email = ?"
			row = db.QueryRow(query,value)
		default:
			logger.Error("Error invalid param in FindUser",errors.New("invalid param in function FindUser"),zap.String("journey","FindUser"))
			return nil,rest_err.NewInternalServerError("database error")
	}

	var retrievedEmail,name,password string
	var age []uint8
	var id uint
	err = row.Scan(&retrievedEmail,&name,&age,&password,&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil,rest_err.NewNotFoundError("User not found")
		}
		logger.Error("Error trying to find user",err,zap.String("journey","FindUserByEmail"))
		return nil,rest_err.NewInternalServerError("Database error")
	}
	user := model.NewUserDomain(
		retrievedEmail,
		password,
		name,
		uint8(age[0]),
		id,
	)
	logger.Info("USER FINDED BY EMAIL")
	return user,nil
}