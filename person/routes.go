package person

import "github.com/gin-gonic/gin"

// Routes Routes
func Routes(route *gin.RouterGroup) {
	route.GET("/:id", getByID)
	route.GET("", get)
	route.POST("", create)
	route.DELETE("/:id", delete)
}
