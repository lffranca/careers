package person

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/careers/model"
	"github.com/lffranca/careers/util"
)

func create(c *gin.Context) {
	var person model.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	db, errDB := c.MustGet(util.ConstContextDB).(*sql.DB)
	if !errDB {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	personNew, errPerson := createPerson(db, &person)
	if errPerson != nil {
		log.Println(errPerson)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, personNew)
}
