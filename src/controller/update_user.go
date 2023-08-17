package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/virussv/api-rest-golang/src/configuration/logger"
	"github.com/virussv/api-rest-golang/src/configuration/validation"
	"github.com/virussv/api-rest-golang/src/controller/model/request"
	"github.com/virussv/api-rest-golang/src/model"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller",
	zap.String("journey","Updateuser"),
)
	var userRequest request.UserUpdateRequest;

	userId := c.Param("userId")
	if err := c.ShouldBindJSON(&userRequest);err != nil || strings.TrimSpace(userId) == ""{
		logger.Error("Error trying to validate user info",err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code,restErr)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)
	err := uc.service.UpdateUser(userId,domain)
	if err != nil {
		logger.Error(
			"Error trying to call UpdateUser service",
			err,
		)
		c.JSON(err.Code,err)
		return
	}

	logger.Info("User created succesfully",
	zap.String("userId",userId),
	zap.String("journey","Updateuser"),
)
	c.Status(http.StatusOK)
}