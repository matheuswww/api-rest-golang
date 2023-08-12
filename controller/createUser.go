package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/virussv/api-rest-golang/configuration/validation"
	"github.com/virussv/api-rest-golang/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest;
	if err := c.ShouldBindJSON(&userRequest);err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code,restErr)
		return
	}
}