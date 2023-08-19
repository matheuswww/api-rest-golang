package repository

import (
	"strconv"

	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(id string) (*rest_err.RestErr) {
	logger.Info("Init DeleteUser repository",zap.String("journey","deleteUser"))
	idRes,err := strconv.ParseUint(id,10,64)
	if err != nil {
		logger.Error("Error trying convert id",err,zap.String("journey","deleteUser"))
		return rest_err.NewBadRequestError("invalid id")
	}
	db := ur.databaseConnection 
	query := "DELETE FROM users WHERE id = ?"
	result,err := db.Exec(query,uint(idRes))
	if err != nil {
		logger.Error("Error trying delete user",err,zap.String("journey","deleteUser"))
		return rest_err.NewInternalServerError("database error")
	}
	if res,_ := result.RowsAffected();res == 0 {
		logger.Error("Error trying delete user",err,zap.String("journey","deleteUser"))
		return rest_err.NewNotFoundError("user not found")
	}
	return nil
}  