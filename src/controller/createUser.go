package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/validation"
	"github.com/virussv/api-rest-golang/src/controller/model/request"
	"github.com/virussv/api-rest-golang/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
	zap.String("journey","createuser"),
)
	var userRequest request.UserRequest;
	if err := c.ShouldBindJSON(&userRequest);err != nil {
		logger.Error("Error trying to validate user info",err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code,restErr)
		return
	}
	response := response.UserResponse  {
		Email: userRequest.Email,
		ID: "Test",
		Name: userRequest.Name,
		Age: userRequest.Age,
	}
	logger.Info("User created succesfully",
	zap.String("journey","createuser"),
)
	c.JSON(http.StatusOK,response)
}