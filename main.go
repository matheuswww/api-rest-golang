package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/controller"
	"github.com/virussv/api-rest-golang/src/controller/routes"
	"github.com/virussv/api-rest-golang/src/model/service"
)

func main() {
	logger.Info("About to start user applicataion")
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRouter(&router.RouterGroup,userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}