package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/blackironj/gin-test/storage"
	"github.com/blackironj/gin-test/util"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var req storage.User
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		req.PID = util.UUIDv4()
		if err := storage.GetInstance().Set(&req); err != nil {
			c.JSON(http.StatusConflict, gin.H{})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"pid": req.PID,
		})
	})

	r.GET("/user/:email", func(c *gin.Context) {
		user := storage.GetInstance().GetByEmail(c.Param("email"))
		c.JSON(http.StatusOK, &user)
	})

	r.DELETE("/user/:email", func(c *gin.Context) {
		storage.GetInstance().DeleteByEmail(c.Param("email"))
		c.JSON(http.StatusOK, gin.H{})
	})

	r.Run(":8080")
}
