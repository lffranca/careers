package person

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/careers/util"
)

// GetByID GetByID
func GetByID(c *gin.Context) {
	db, errDB := c.MustGet(util.ConstContextDB).(*sql.DB)
	if !errDB {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	if err := db.PingContext(context.Background()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"OK": "OK",
	})
}
