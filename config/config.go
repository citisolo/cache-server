package config

import "github.com/JSainsburyPLC/third-party-token-server/db"


type Config struct {
	CacheManager db.CacheManager 
}

func GetConfig() *Config {
	return &Config{
		CacheManager: db.CacheManager{},
	}
}