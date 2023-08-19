package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/validation"
	"github.com/virussv/api-rest-golang/src/controller/model/request"
	"github.com/virussv/api-rest-golang/src/model"
	"github.com/virussv/api-rest-golang/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller",
	zap.String("journey","loginUser"),
)
	var userRequest request.UserLogin;

	if err := c.ShouldBindJSON(&userRequest);err != nil {
		logger.Error("Error trying to validate user info",err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code,restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
		0,
	)
	domainResult,token,err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error(
			"Error trying to call loginUser service",
			err,
		)
		c.JSON(err.Code,err)
		return
	}

	logger.Info("User created succesfully",
	zap.String("userId",domainResult.GetEmail()),
	zap.String("journey","loginUser"),
	)

	c.Header("Authorization",token)

	c.JSON(http.StatusOK,view.ConvertDomainToResponse(
		domainResult,
	))
}