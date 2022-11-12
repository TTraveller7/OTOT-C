package common

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var allowOrigins []string = []string{
	"*",
}

func GetUrl(envKey string, defaultAddr string) string {
	if envUrl := os.Getenv(envKey); len(envUrl) != 0 {
		return envUrl
	}
	return defaultAddr
}

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = allowOrigins
	return cors.New(config)
}
