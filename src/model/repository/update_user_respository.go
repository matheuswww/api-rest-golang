package repository

import (
	"github.com/matheuswww/api-rest-golang/src/configuration/logger"
	"github.com/matheuswww/api-rest-golang/src/configuration/rest_err"
	"github.com/matheuswww/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(userId string,userDomain model.UserDomainInterface,
	) (
	*rest_err.RestErr) {
	logger.Info("Init UpdateUser repository",zap.String("journey","UpdateUser"))
	db := ur.databaseConnection
	defer db.Close()
	query := `UPDATE users SET 
			name = CASE WHEN ? <> '' THEN ? ELSE name END,
			age = CASE WHEN ? > 0 THEN ? ELSE age END WHERE id = ?`
	result,err := db.Exec(query,
		userDomain.GetName(),
		userDomain.GetName(),
		userDomain.GetAge(),
		userDomain.GetAge(),
		userId,
	)
	if err != nil {
		logger.Error("Error trying to update user",err,zap.String("journey","UpdateUser"))
		return rest_err.NewInternalServerError("database error")
	}
	id,err := result.LastInsertId()
	if err != nil {
		logger.Error("Error trying to getting id",err,zap.String("journey","UpdateUser"))
		return rest_err.NewInternalServerError("database error")
	}
	userDomain.SetId(uint(id))
	logger.Info("USER INSERTED IN MYSQL DATABASE")
	return nil
}