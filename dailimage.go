package main

import (
	"fmt"
	"local/eldritchgarden/dailimage/config"
	"local/eldritchgarden/dailimage/image"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	if len(config.ENV.TrustedProxies) > 0 {
		router.SetTrustedProxies(config.ENV.TrustedProxies)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	router.GET("/random", image.GetRandomImage)
	router.GET("/random/*subpath", image.GetRandomImagePath)

	router.Run(fmt.Sprintf("%s:%s", config.ENV.Addr, config.ENV.Port))
}
