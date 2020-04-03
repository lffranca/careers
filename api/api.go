package api

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	dbconn "github.com/lffranca/careers/db_conn"
	"github.com/lffranca/careers/person"
)

// InitAPI InitAPI
func InitAPI(wgParent *sync.WaitGroup) {
	defer wgParent.Done()

	dbConn, errDBConn := dbconn.NewDBConn()
	if errDBConn != nil {
		log.Fatalln(errDBConn)
	}

	router := gin.Default()

	routerAPI := router.Group("/api")
	routerAPI.Use(dbConn.MiddlewareDB())
	{
		routerAPI.GET("/:id", person.GetByID)
	}
}
