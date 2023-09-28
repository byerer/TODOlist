package controller

import (
	"TODOlist/dao/mysql"
	"TODOlist/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddToDo(c *gin.Context) {
	var todo models.TODO
	c.BindJSON(&todo)
	mysql.DB.Create(todo)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"todos":  todo,
	})
}

func DeleteToDo(c *gin.Context) {
	index, _ := strconv.Atoi(c.Param("index"))
	err := mysql.DB.Where("id=?", index).Delete(&models.TODO{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "record does not exist",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func UpdateToDo(c *gin.Context) {
	index, _ := strconv.Atoi(c.Param("index"))
	var todo models.TODO
	c.BindJSON(&todo)
	mysql.DB.Model(&models.TODO{}).Where("id=?", index).Updates(todo)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func GetAllToDO(c *gin.Context) {
	var todo []models.TODO
	mysql.DB.Find(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"obj":    todo,
	})
}

func GetTodo(c *gin.Context) {
	index, _ := strconv.Atoi(c.Param("index"))
	var todo models.TODO
	todo.ID = int64(index)
	mysql.DB.First(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"obj":    todo,
	})
}
