package main

import (
	"database/sql"

	"github.com/virussv/api-rest-golang/src/controller"
	"github.com/virussv/api-rest-golang/src/model/repository"
	"github.com/virussv/api-rest-golang/src/model/service"
)

func initDependencies(db *sql.DB) (controller.UserControllerInterface){
	repo := repository.NewUserRepository(db)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}