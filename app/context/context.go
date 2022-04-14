package context

import (
	"github.com/JSainsburyPLC/third-party-token-server/app/logging"
	"github.com/JSainsburyPLC/third-party-token-server/db"
)


type AppContext struct {
	Cache db.Cache
	Logger logging.AppLogs
}