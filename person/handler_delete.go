package person

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/careers/util"
)

func delete(c *gin.Context) {
	db, errDB := c.MustGet(util.ConstContextDB).(*sql.DB)
	if !errDB {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	id, errToInt := strconv.Atoi(c.Param("id"))
	if errToInt != nil {
		log.Println(errToInt)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid param"})
		return
	}

	errPerson := deletePersonByID(db, id)
	if errPerson != nil {
		log.Println(errPerson)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}
