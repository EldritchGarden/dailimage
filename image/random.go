package image

import (
	"errors"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/eldritchgarden/dailimage/internal/config"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var imageCache map[string]bool
var imageCacheSize int

// Called by filepath.WalkDir to cache an image file path
func cacheFile(path string, info fs.FileInfo, err error) error {
	if err != nil {
		log.Trace("Error accessing path ", path, ": ", err)
		return err
	}
	if !info.IsDir() {
		log.Trace("Cached ", path)
		imageCache[path] = true
	}
	return nil
}

func buildImageCache() error {
	log.Info("Building file cache...")
	imageCache = make(map[string]bool)
	err := filepath.Walk(config.ENV.MediaRoot, cacheFile)
	if err != nil {
		return err
	}
	imageCacheSize = len(imageCache)
	return nil
}

// Returns a slice of image paths from the cache that are under the specified root directory.
// If the cache is low or empty, it attempts to rebuild it. Allows one retry to rebuild the cache if no files are found.
func sliceImageCache(root string, try int) ([]string, error) {
	log.Debugf("(Try %d) Slicing image cache for root: %s", try, root)

	minSize := int(config.ENV.UniqueFactor * float32(imageCacheSize))
	if minSize < 1 {
		minSize = 1
	}
	log.Debugf("Current cache size: %d - Rebuild threshold: %d",
		len(imageCache), minSize)

	if try > 0 || len(imageCache) < minSize {
		err := buildImageCache()
		if err != nil {
			return []string{}, err
		}
	}

	// only get files under the specified root
	files := make([]string, 0, len(imageCache))
	i := 0
	for path := range imageCache {
		if strings.HasPrefix(path, root) {
			files = append(files, path)
			i++
		}
	}

	// the slice may be empty even if the cache is not, depending on root
	// if so, try rebuilding the cache once before giving up
	if len(files) == 0 && try < 1 {
		log.Debug("Rebuilding cache and retrying slice...")
		return sliceImageCache(root, try+1)
	} else if len(files) == 0 {
		return []string{}, errors.New("no images found in directory: " + root)
	}

	return files, nil
}

// Returns a random image path from the given root directory or its subdirectories
func randomImagePath(root string) (string, error) {
	log.Debug("Getting random image from root: ", root)
	files, err := sliceImageCache(root, 0)
	if err != nil {
		return "", err
	}

	// Select a random image from the cache
	path := files[rand.Intn(len(files))]
	if config.ENV.UniqueFactor < 1 {
		// remove from cache to avoid repeats unless factor is 1 to avoid
		// excessive cache rebuilding
		delete(imageCache, path)
	}

	return path, nil
}

func GetRandomImage(c *gin.Context) {
	file, err := randomImagePath(config.ENV.MediaRoot)
	if err != nil {
		c.String(500, "Error retrieving random image: %v", err)
		return
	}

	c.Header("Cache-Control", "no-store") // prevent caching the image

	log.Debug("Serving: ", file)
	c.File(file)
}

// Takes a subdirectory path and returns a random image from that directory
// or its subdirectories. The subpath is relative to the configured media root.
func GetRandomImagePath(c *gin.Context) {
	subpath := c.Param("subpath")
	root := filepath.Join(config.ENV.MediaRoot, subpath)
	log.Debug("Getting random image from subpath: ", subpath, " (root: ", root, ")")
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

	c.Header("Cache-Control", "no-store") // prevent caching the image

	log.Debug("Serving: ", file)
	c.File(file)
}
