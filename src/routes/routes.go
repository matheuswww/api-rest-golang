package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/virussv/api-rest-golang/src/controller"
)

func InitRouter(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId",controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail",controller.FindUserByEmail)
	r.POST("/CreateUser",controller.CreateUser)
	r.PUT("/updateUser/:userId",controller.UpdateUser)
	r.DELETE("/deleteUser/:userId",controller.DeleteUser)
}