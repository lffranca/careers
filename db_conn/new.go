package db_conn

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/careers/util"
)

// DB DB
type DB struct {
	DB *sql.DB
}

// NewDBConn NewDBConn
func NewDBConn() (*DB, error) {
	db, errDB := createConn()
	if errDB != nil {
		return nil, errDB
	}

	return &DB{DB: db}, nil
}

// MiddlewareDB MiddlewareDB
func (conn *DB) MiddlewareDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := conn.DB.PingContext(context.Background()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Try again"})
			c.Abort()
			return
		}

		c.Set(util.ConstContextDB, conn.DB)
		c.Next()
	}
}
