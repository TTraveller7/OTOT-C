package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	c.Status(http.StatusAccepted)
}

func EngineWithHealthcheck() *gin.Engine {
	engine := gin.New()
	engine.Use(
		gin.LoggerWithConfig(
			gin.LoggerConfig{
				SkipPaths: []string{HEALTHCHECK_URL},
			},
		),
		gin.Recovery(),
	)
	engine.GET(HEALTHCHECK_URL, Healthcheck)

	return engine
}
