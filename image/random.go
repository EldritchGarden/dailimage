package image

import (
	"local/eldritchgarden/dailimage/config"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
)

func RandomImagePath(c *gin.Context) {
	root := config.ENV.MediaRoot
	for {
		file, err := os.Open(root)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		list, err := file.Readdirnames(-1)
		if err != nil {
			panic(err)
		}

		randIndex := rand.Intn(len(list))
		result := list[randIndex]
		if f, e := os.Lstat(root + "/" + result); f.IsDir() && e == nil {
			root = root + "/" + result // recurse into directory
		} else if e != nil {
			panic(e)
		} else {
			c.File(root + "/" + result)
			return
		}
	}
}
