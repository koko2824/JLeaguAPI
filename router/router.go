package router

import (
	"github.com/gin-gonic/gin"
	"os"
	"scraping/controller"
)

func Init() {
	r := router()
	r.Run(":" + os.Getenv("PORT"))
}

func router() *gin.Engine {
	r := gin.Default()
	ctrl := controller.Controller{}
	r.GET("/", ctrl.Ranking)
	r.GET("/difference", ctrl.ScoreDifference)
	return r
}
