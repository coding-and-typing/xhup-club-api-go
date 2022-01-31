package router

import (
	"os"

	"github.com/coding-and-typing/xhup-club-api-go/controller"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	engine := gin.New()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)
	// Add a logger middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	engine.Use(logger.SetLogger(), gin.Recovery())
	registerBaseAPI(engine)
	return engine
}

// registerBaseAPI ...
func registerBaseAPI(engine *gin.Engine) {
	engine.GET("/health", controller.Health)
}
