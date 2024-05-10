package config

import (
	"github.com/depender/email-sequence-service/constants"
	"github.com/gin-gonic/gin"
)

// Middleware ensures that config will remain same for a single request in his whole journey
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		app := getApplication()
		if app == nil {
			c.Abort()
		}

		conf := app.watcher.GetConfig()
		c.Set(constants.Config, conf)
		c.Next()
	}
}
