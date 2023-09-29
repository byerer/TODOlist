package main

import (
	"TODOlist/controller"
	"TODOlist/dao/mysql"
	"TODOlist/middlewares/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//gorm
	err := mysql.InitMysql()
	if err != nil {
		fmt.Println(err)
	}

	//router
	r := gin.Default()
	//user
	r.GET("/login", controller.Login)
	r.POST("/register", controller.Register)

	//todo
	todo := r.Group("/todo", jwt.MiddleWareJWT)
	{
		todo.GET("", controller.GetAllToDO)
		todo.POST("", controller.AddToDo)
		todo.DELETE("/:id", controller.VerifyPermission, controller.DeleteToDo)
		todo.PUT("/:id", controller.VerifyPermission, controller.UpdateToDo)
		todo.GET("/:id", controller.VerifyPermission, controller.GetTodo)
	}
	_ = r.Run(":8080")

}
