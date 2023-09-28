package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TODO struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

var todos []TODO

func main() {
	r := gin.Default()
	r.POST("/todo", func(c *gin.Context) {
		var todo TODO
		c.BindJSON(&todo)
		todos = append(todos, todo)
		fmt.Println(todos)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.DELETE("/todo/:index", func(c *gin.Context) {
		index, _ := strconv.Atoi(c.Param("index"))
		todos = append(todos[:index], todos[index+1:]...)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.PUT("todo/:index", func(c *gin.Context) {
		index, _ := strconv.Atoi(c.Param("index"))
		var todo TODO
		c.BindJSON(&todo)
		todos[index] = todo
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.GET("/todo", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})

	r.GET("todo/:index", func(c *gin.Context) {
		index, _ := strconv.Atoi(c.Param("index"))
		c.JSON(http.StatusOK, todos[index])
	})
	r.Run(":8080")
}
