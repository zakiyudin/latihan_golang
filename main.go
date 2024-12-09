package main

import "github.com/gin-gonic/gin"

func main() {
	// inisialisasi router
	router := gin.Default()

	// Middleware: Logger
	router.Use(gin.Logger())

	// Middleware: Recovery
	router.Use(gin.Recovery())

	// basic endpoints
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	// endpoints parameters
	router.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(200, gin.H{
			"message": "hello " + name + "!",
		})
	})

	// endpoint login (POST)
	router.POST("/login", func(ctx *gin.Context) {
		var loginData struct {
			Email    string `json:email`
			Password string `json:password`
		}

		if err := ctx.ShouldBindJSON(&loginData); err != nil {
			ctx.JSON(401, gin.H{
				"message": "invalid body request",
			})
			return
		}

		if loginData.Email == "admin@mail.com" && loginData.Password == "12345789" {
			ctx.JSON(200, gin.H{
				"message": "success",
			})
		} else {
			ctx.JSON(401, gin.H{
				"message": "error credentials",
			})
		}
	})

	// endpoints query parameter
	router.GET("/user", func(ctx *gin.Context) {
		name := ctx.Query("name")

		if name == "" {
			ctx.JSON(400, gin.H{
				"message": "invalid query",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "hallo " + name + "!!!",
		})
	})

	router.Run(":8091")
}
