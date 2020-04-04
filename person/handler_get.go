package person

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/careers/util"
)

func get(c *gin.Context) {
	db, errDB := c.MustGet(util.ConstContextDB).(*sql.DB)
	if !errDB {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	search := c.Query("search")
	hero := c.Query("hero")

	people, errPeople := getPeople(db, search, hero)
	if errPeople != nil {
		log.Println(errPeople)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, people)
}
