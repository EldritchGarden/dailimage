package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"local/eldritchgarden/dailimage/config"
	"local/eldritchgarden/dailimage/image"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates
var templateFS embed.FS

//go:embed static
var staticFS embed.FS

func main() {
	router := gin.Default()
	if len(config.ENV.TrustedProxies) > 0 {
		router.SetTrustedProxies(config.ENV.TrustedProxies)
	} else {
		router.SetTrustedProxies(nil) // Disable trusted proxies if none are set
	}

	// Load embedded templates
	router.SetHTMLTemplate(template.Must(template.New("").ParseFS(templateFS,
		"templates/*")))

	staticFS, _ := fs.Sub(staticFS, "static")
	router.StaticFS("/static", http.FS(staticFS))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	router.GET("/random", image.GetRandomImage)
	router.GET("/random/*subpath", image.GetRandomImagePath)
	router.GET("/slideshow", image.GetSlideshow)

	router.Run(fmt.Sprintf("%s:%s", config.ENV.Addr, config.ENV.Port))
}
