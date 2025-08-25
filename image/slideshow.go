package image

import (
	"html/template"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetSlideshow(c *gin.Context) {
	interval := c.DefaultQuery("interval", "30") // default to 30s
	subpath := c.DefaultQuery("subpath", "")
	mode := strings.ToLower(c.DefaultQuery("mode", "frame"))

	if i, e := strconv.Atoi(interval); e != nil || i < 1 {
		interval = "30" // reset to default if invalid
	}
	if mode != "frame" && mode != "full" {
		mode = "frame" // reset to default if invalid
	}

	source := "/random"
	if subpath != "" {
		source += "/" + subpath
	}

	c.HTML(200, "slideshow.tmpl", gin.H{
		"source":   template.HTMLEscapeString(source),
		"interval": template.HTMLEscapeString(interval),
		"style":    template.HTMLEscapeString("/static/img-" + mode + ".css"),
	})
}
