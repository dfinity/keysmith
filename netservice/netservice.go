package netservice

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const NETWORK_CMD = "netservice"

const NETWORK_CMD_START = "start"

const NETWORK_CMD_STOP = "stop"

const NETWORK_CMD_RESTART = "restart"

var router = gin.Default()

func start() {

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	router.Run(":8081") // listen and serve on 0.0.0.0:8081
}
