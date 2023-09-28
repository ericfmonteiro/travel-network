package routes

import (
	"fmt"

	"github.com/ericfmonteiro/travel-network/app/handler"
	"github.com/gin-gonic/gin"
)

func InitGin() {
	router := gin.Default()

	// CORS stuff
	router.Use(corsMiddleware())
	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(204)
	})

	// App default
	router.GET("/", handler.Login)

	// Users
	router.GET("/users/name/:name", handler.GetUserByName)
	router.POST("/users/create", handler.CreateUser)
	router.PUT("/users/:username/edit", handler.EditUser)
	router.GET("/users/:username/connected-posts", handler.GetConnectedPosts)

	// Posts
	router.POST("/posts/create", handler.CreatePost)
	router.GET("/posts/likes/:postid", handler.GetLikesByPost)
	router.POST("/posts/likes", handler.CreateLike)

	// Comments
	router.POST("/posts/comments", handler.CreateComment)
	router.GET("/posts/comments/:postid", handler.GetCommentsByPost)

	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Could not start server")
		panic(err)
	}
	fmt.Println("Server started on :8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, Accept-Language, Content-Language")

		c.Next()
	}
}
