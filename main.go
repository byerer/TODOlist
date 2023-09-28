package main

import (
	"TODOlist/controller"
	"TODOlist/dao/mysql"
	"TODOlist/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

var todos []models.TODO

func main() {
	//gorm
	err := mysql.InitMysql()
	if err != nil {
		fmt.Println(err)
	}

	//router
	r := gin.Default()
	r.POST("/todo", controller.AddToDo)
	r.DELETE("/todo/:index", controller.DeleteToDo)
	r.PUT("/todo/:index", controller.UpdateToDo)
	r.GET("/todo", controller.GetAllToDO)
	r.GET("/todo/:index", controller.GetTodo)
	r.Run(":8080")

}
