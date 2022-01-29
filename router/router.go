package router

import (
	"github.com/coding-and-typing/xhup-club-api-go/controller"

	"github.com/gin-gonic/gin"
)


// InitRouter ...
func InitRouter() *gin.Engine {
	engine := gin.Default()
	registerBaseAPI(engine)
	return engine
}

// registerBaseAPI ...
func registerBaseAPI(engine *gin.Engine) {
	engine.GET("/health", controller.Health)
}