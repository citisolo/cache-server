package config

import "github.com/JSainsburyPLC/third-party-token-server/db"


type Config struct {
	CacheManager db.CacheManager 
	KeyFile  string
	CertFile string
}

func GetConfig() *Config {
	return &Config{
		CacheManager: db.CacheManager{},
		CertFile: "localhost.crt",
		KeyFile: "localhost.key",
	}
}