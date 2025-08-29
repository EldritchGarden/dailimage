package config

import (
	"strconv"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Addr           string   // Bind address
	Port           string   // Bind port
	MediaRoot      string   // Directory for image files
	TrustedProxies []string // Optional trusted proxies
	LogLevel       log.Level
	UniqueFactor   float32 // Threshold before rebuilding cache
}

func setLogLevel(level string) log.Level {
	logLevel, err := log.ParseLevel(strings.ToLower(level))
	if err != nil {
		log.Warnf("Invalid LOG_LEVEL '%s', defaulting to 'info'", level)
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	log.Info("Log level: ", level)
	return logLevel
}

func FromEnv() Config {
	level, ok := syscall.Getenv("LOG_LEVEL")
	if !ok {
		level = "info"
	}
	logLevel := setLogLevel(level)

	addr, ok := syscall.Getenv("BIND_ADDR")
	if !ok {
		addr = "0.0.0.0"
	}
	log.Debug("Bind address: ", addr)

	port, ok := syscall.Getenv("BIND_PORT")
	if !ok {
		port = "8080"
	}
	log.Debug("Bind port: ", port)

	mediaRoot, ok := syscall.Getenv("MEDIA_ROOT")
	if !ok {
		panic("MEDIA_ROOT environment variable is required")
	}
	log.Debug("Media root: ", mediaRoot)

	var trustedProxies []string
	if trustedProxiesSetting, ok := syscall.Getenv("TRUSTED_PROXIES"); ok {
		trustedProxies = strings.Split(trustedProxiesSetting, ",")
	}
	log.Debug("Trusted proxies: ", trustedProxies)

	var uniqueFactor float32 = 0.5
	uniqueFactorSetting, ok := syscall.Getenv("UNIQUE_FACTOR")
	if !ok {
		uniqueFactorSetting = "0.5"
	}

	if uf, err := strconv.ParseFloat(uniqueFactorSetting, 32); err == nil {
		if uf >= 0.01 && uf <= 1 {
			uniqueFactor = float32(uf)
		} else {
			log.Warnf("UNIQUE_FACTOR '%s' out of range, default to 0.5",
				uniqueFactorSetting)
		}
	} else {
		log.Warnf("Invalid UNIQUE_FACTOR '%s', default to 0.5",
			uniqueFactorSetting)
	}
	log.Debug("Unique factor: ", uniqueFactor)

	// if set to 1, the cache will never refresh and images
	// may repeat frequently
	if uniqueFactor == 1 {
		log.Warn("UNIQUE_FACTOR is set to 1, cache refreshes disabled")
	}

	return Config{
		Addr:           addr,
		Port:           port,
		MediaRoot:      mediaRoot,
		TrustedProxies: trustedProxies,
		LogLevel:       logLevel,
		UniqueFactor:   uniqueFactor,
	}
}

// Expose this for easy access
var ENV Config = FromEnv()
