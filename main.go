package main

import (
	"TODOlist/controller"
	"TODOlist/dao/mysql"
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
	r.POST("/todo", controller.AddToDo)
	r.DELETE("/todo/:index", controller.DeleteToDo)
	r.PUT("/todo/:index", controller.UpdateToDo)
	r.GET("/todo", controller.GetAllToDO)
	r.GET("/todo/:index", controller.GetTodo)
	r.GET("/parsetoken", controller.ParseToken)
	_ = r.Run(":8080")

}
