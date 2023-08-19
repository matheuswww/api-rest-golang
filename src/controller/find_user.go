package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/rest_err"
	"github.com/virussv/api-rest-golang/src/model"
	"github.com/virussv/api-rest-golang/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUser(c *gin.Context) {
	logger.Info("Init findUser controller",
		zap.String("journey", "findUser"),
	)
	var val string
	var domainResult model.UserDomainInterface
	var err *rest_err.RestErr
	if val = c.Param("userId"); val != "" {
		domainResult, err = uc.service.FindUser("id", val)
	} else if val = c.Param("userEmail"); val != "" {
		domainResult, err = uc.service.FindUser("email", val)
	}
	if err != nil {
		logger.Error("Error trying find user", err)
		c.JSON(err.Code, err)
		return
	}
	logger.Info("Finded use succesfully",
		zap.String("userId", domainResult.GetEmail()),
		zap.String("journey", "findUser"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
