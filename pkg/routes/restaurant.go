package routes

import (
	"github.com/abhay676/today-menu/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func LoadRestaurantRouter(route *gin.Engine) {

	r := route.Group("/restaurant")

	r.POST("/onboarding", controllers.Onboarding)
	// r.GET("/detail")
	// r.PUT("/update")
	// r.DELETE("/remove")
}
