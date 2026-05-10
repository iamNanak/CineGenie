package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controllers "github.com/iamNanak/CineGenie/Server/StreamServer/controllers"
)

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("id: ", id, "work done")
// }

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

	// numWorkers := 10

	// var wg sync.WaitGroup

	// for i := 0; i < numWorkers; i++ {
	// 	wg.Add(1)
	// 	go worker(i, &wg)
	// }

	// wg.Wait()
}
