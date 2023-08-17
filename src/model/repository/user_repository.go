package repository

import (
	"database/sql"

	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"github.com/virussv/api-rest-golang/src/model"
)

func NewUserRepository(database *sql.DB) UserRepository{
		return &userRepository {
			database,
		}
}

type userRepository struct {
	databaseConnection *sql.DB
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface,*rest_err.RestErr)
	FindUser(string,string) (model.UserDomainInterface,*rest_err.RestErr)
	DeleteUser(string) (*rest_err.RestErr)
	UpdateUser(string,model.UserDomainInterface) *rest_err.RestErr
}