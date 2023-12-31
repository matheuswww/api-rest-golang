package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/api-rest-golang/src/configuration/logger"
	"github.com/matheuswww/api-rest-golang/src/configuration/validation"
	"github.com/matheuswww/api-rest-golang/src/controller/model/request"
	"github.com/matheuswww/api-rest-golang/src/model"
	"github.com/matheuswww/api-rest-golang/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
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

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
		0,
	)
	domainResult,err := uc.service.CreateUser(domain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
		)
		c.JSON(err.Code,err)
		return
	}

	logger.Info("User created succesfully",
	zap.String("userId",domainResult.GetEmail()),
	zap.String("journey","createuser"),
)
	c.JSON(http.StatusOK,view.ConvertDomainToResponse(
		domainResult,
	))
}