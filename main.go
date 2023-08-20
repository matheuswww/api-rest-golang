package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matheuswww/api-rest-golang/src/configuration/database/mysql"
	"github.com/matheuswww/api-rest-golang/src/configuration/logger"
	"github.com/matheuswww/api-rest-golang/src/controller/routes"
)

func main() {
	logger.Info("About to start user applicataion")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := mysql.NewMysqlConnection()
	if err != nil {
		log.Fatal("Error to connect mysql db")
	}

	userController := initDependencies(db)

	router := gin.Default()
	routes.InitRouter(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error to load router", err)
		log.Fatal(err)
	}
}