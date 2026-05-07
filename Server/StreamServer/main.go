package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controllers "github.com/iamNanak/CineGenie/Server/StreamServer/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, ("hello CineGenie"))
	})

	router.GET("/movies", controllers.GetMovies())

	router.GET("/movie/:imdb_id", controllers.GetMovieByID())

	router.POST("/addmovie", controllers.AddMovie())

	router.POST("/register", controllers.RegisterUser())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}

}
