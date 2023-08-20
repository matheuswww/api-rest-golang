package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/matheuswww/api-rest-golang/src/configuration/logger"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller",zap.String("journey","deleteUser"))
	id := c.Param("userId")
	err := uc.service.DeleteUser(id)
	if err != nil {
		logger.Error("Error trying deleteUser",err,zap.String("journey","deleteUser"))
		c.JSON(err.Code,err)
		return
	}
	c.Status(http.StatusOK)
}