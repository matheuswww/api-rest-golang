package repository

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/matheuswww/api-rest-golang/src/configuration/logger"
	"github.com/matheuswww/api-rest-golang/src/configuration/rest_err"
	"github.com/matheuswww/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUser(queryType string,value string) (model.UserDomainInterface,*rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository",zap.String("journey","FindUserByEmail"))

	db := ur.databaseConnection
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
			return nil,rest_err.NewBadRequestError("invalid param")
	}

	var retrievedEmail,name,password string
	var age []uint8
	var id uint
	err := row.Scan(&retrievedEmail,&password,&name,&age,&id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("User not found",err,zap.String("journey","FindUser"))
			return nil,rest_err.NewNotFoundError("User not found")
		}
		logger.Error("Error trying to find user",err,zap.String("journey","FindUser"))
		return nil,rest_err.NewInternalServerError("Database error")
	}
	ageValue, err := strconv.Atoi(string(age))
	if err != nil {
			logger.Error("Error converting age to int", err, zap.String("journey", "FindUser"))
			return nil, rest_err.NewInternalServerError("Database error")
	}
	user := model.NewUserDomain(
		retrievedEmail,
		password,
		name,
		uint8(ageValue),
		id,
	)
	logger.Info("USER FINDED",zap.String("journey","FindUser"))
	return user,nil
}


func (ur *userRepository) FindUserByEmailAndPassword(email,password string) (model.UserDomainInterface,*rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword repository",zap.String("journey","findUserByEmailAndPassword"))
	db := ur.databaseConnection
	var retrievedEmail,retrievedPassword,name string
	var id uint 
	var age uint8
	query := "SELECT email,password,name,age,id FROM users WHERE email = ? AND password = ?"
	row := db.QueryRow(query,email,password)
	err := row.Scan(&retrievedEmail,&retrievedPassword,&name,&age,&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil,rest_err.NewNotFoundError("User not found")
		}
		logger.Error("Error trying to find user",err,zap.String("journey","findUserByEmailAndPassword"))
		return nil,rest_err.NewInternalServerError("Database error")
	}
	user := model.NewUserDomain(
		retrievedEmail,
		retrievedPassword,
		name,
		age,
		id,
	)
	logger.Info("USER FINDED BY EMAIL AND PASSWORD",zap.String("journey","FindUser"))
	return user,nil
}