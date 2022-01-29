package main

import (
	"github.com/coding-and-typing/xhup-club-api-go/router"
)

func main() {
	engine := router.InitRouter()
	engine.Run(":8080")
}
