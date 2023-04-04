package main

import (
	controllers "golang-mongodb/controllers"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
)

func main() {
	router := gin.New()

	router.GET("/users", controllers.GetUsers)

	router.GET("/user/create", controllers.CreateUser)

	router.Run(":3000")
}
