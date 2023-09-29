package controller

import (
	"TODOlist/dao/mysql"
	"TODOlist/middlewares/jwt"
	"TODOlist/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var user models.User
	username := c.Query("username")
	password := c.Query("password")
	mysql.DB.Where("username = ? AND password = ?", username, password).Find(&user)
	token, _ := jwt.GenerateToken(user.UserID, user.Username)
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"user":    user,
		"token":   token,
	})
}

func Register(c *gin.Context) {
	var user models.User
	_ = c.BindJSON(&user)
	mysql.DB.Create(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
		"user":    user,
	})
}

func ParseToken(c *gin.Context) {
	token := c.Query("token")
	claims, err := jwt.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "login expired",
		})
	}
	fmt.Println(claims)
	c.JSON(http.StatusOK, gin.H{
		"userid":   claims.UserID,
		"username": claims.Username,
	})
}
