package image

import (
	"errors"
	"io/fs"
	"local/eldritchgarden/dailimage/config"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
)

func randomImagePath(root string) (string, error) {
	for {
		file, err := os.Open(root)
		if err != nil {
			return "", err
		}
		defer file.Close()

		list, err := file.Readdirnames(-1)
		if err != nil {
			return "", err
		} else if len(list) == 0 {
			return "", errors.New("no files found in directory: " + root)
		}

		randIndex := rand.Intn(len(list))
		result := list[randIndex]
		if f, e := os.Lstat(root + "/" + result); f.IsDir() && e == nil {
			root = root + "/" + result // recurse into directory
		} else if e != nil {
			return "", err
		} else {
			return root + "/" + result, nil
		}
	}
}

func GetRandomImage(c *gin.Context) {
	file, err := randomImagePath(config.ENV.MediaRoot)
	if err != nil {
		c.String(500, "Error retrieving random image: %v", err)
		return
	}
	c.File(file)
}

// Takes a subdirectory path and returns a random image from that directory
// or its subdirectories. The subpath is relative to the configured media root.
func GetRandomImagePath(c *gin.Context) {
	subpath := c.Param("subpath")
	root := config.ENV.MediaRoot + "/" + subpath
	if _, err := os.Stat(root); errors.Is(err, fs.ErrNotExist) {
		c.String(404, "Subdirectory not found: %s", subpath)
		return
	} else if err != nil {
		c.String(500, "Error accessing subdirectory: %v", err)
		return
	}

	file, err := randomImagePath(root)
	if err != nil {
		c.String(500, "Error retrieving image: %v", err)
		return
	}
	c.File(file)
}
