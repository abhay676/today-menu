package routes

import "github.com/gin-gonic/gin"

func LoadRouter(route *gin.Engine) {

	LoadRestaurantRouter(route)
}
