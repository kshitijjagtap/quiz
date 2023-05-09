package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kshitijjagtap/quiz_usingreact/controllers"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/api/answer", controllers.Answer_f)
	router.POST("/api/submit", controllers.Submit)
	router.Run(":9000")
}
