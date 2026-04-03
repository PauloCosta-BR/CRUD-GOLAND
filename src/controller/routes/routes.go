package routes

import (
	"github.com/PauloCosta-BR/CRUD-GOLAND/src/controller"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	router.GET("/getUserById/:userId", controller.FindUserById)
	router.GET("/getUserByEmail/:userId", controller.FindUserByEmail)
	router.POST("/CreateUser", controller.CreateUser)
	router.PUT("/UpdateUser/:userId", controller.UpdateUser)
	router.DELETE("/DeleteUser/:userId", controller.DeleteUser)

}
