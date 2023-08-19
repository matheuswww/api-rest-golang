package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/virussv/api-rest-golang/src/controller"
	"github.com/virussv/api-rest-golang/src/model"
)

func InitRouter(r *gin.RouterGroup,userController controller.UserControllerInterface) {
	r.GET("/getUserById/:userId",model.VerifyTokenMiddleware,userController.FindUser)
	r.GET("/getUserByEmail/:userEmail",model.VerifyTokenMiddleware,userController.FindUser)
	r.POST("/createUser",model.VerifyTokenMiddleware,userController.CreateUser)
	r.PUT("/updateUser/:userId",model.VerifyTokenMiddleware,userController.UpdateUser)
	r.DELETE("/deleteUser/:userId",model.VerifyTokenMiddleware,userController.DeleteUser)	 

	r.POST("/login",userController.LoginUser)
}