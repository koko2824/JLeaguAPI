package router

import (
	"github.com/gin-gonic/gin"
	"scraping/controller"
)

func Init() {
	r := router()
	//r.Run(":" + os.Getenv("PORT"))
	r.Run("localhost:8080")
}

func router() *gin.Engine {
	r := gin.Default()
	ctrl := controller.Controller{}
	r.GET("/", ctrl.Ranking)
	return r
}
