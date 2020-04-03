package person

import "github.com/gin-gonic/gin"

// Routes Routes
func Routes(route *gin.RouterGroup) {
	route.GET("/:id", getByID)
	route.POST("", create)
}
