package config

import (
	"strings"
	"syscall"
)

type Config struct {
	Addr           string   // Bind address
	Port           string   // Bind port
	MediaRoot      string   // Directory for image files
	TrustedProxies []string // optional trusted proxies
}

func FromEnv() Config {
	addr, ok := syscall.Getenv("BIND_ADDR")
	if !ok {
		addr = "0.0.0.0"
	}

	port, ok := syscall.Getenv("BIND_PORT")
	if !ok {
		port = "8080"
	}

	mediaRoot, ok := syscall.Getenv("MEDIA_ROOT")
	if !ok {
		panic("MEDIA_ROOT environment variable is required")
	}

	var trustedProxies []string
	if trusted, ok := syscall.Getenv("TRUSTED_PROXIES"); ok {
		trustedProxies = strings.Split(trusted, ",")
	}

	return Config{
		Addr:           addr,
		Port:           port,
		MediaRoot:      mediaRoot,
		TrustedProxies: trustedProxies,
	}
}

var ENV Config = FromEnv()
