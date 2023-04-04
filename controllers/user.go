package controllers

import (
	dbhelper "golang-mongodb/common/db"

	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

var userColleciton *mongo.Collection = dbhelper.Client.Database("user_manager").Collection("user")

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, 2)
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, 3)
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, 4)
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, 5)
}
