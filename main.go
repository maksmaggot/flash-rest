package main

import (
	"flash-rest/controllers"
	"flash-rest/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := models.SetupModels()
	r.Use(func(context *gin.Context) {
		context.Set("db", db)
		context.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)


	r.Run()
}
