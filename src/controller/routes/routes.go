package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/virussv/api-rest-golang/src/controller"
)

func InitRouter(r *gin.RouterGroup,userController controller.UserControllerInterface) {
	r.GET("/getUserById/:userId",userController.FindUser)
	r.GET("/getUserByEmail/:userEmail",userController.FindUser)
	r.POST("/createUser",userController.CreateUser)
	r.PUT("/updateUser/:userId",userController.UpdateUser)
	r.DELETE("/deleteUser/:userId",userController.DeleteUser)	 

	r.POST("/login",userController.LoginUser)
}