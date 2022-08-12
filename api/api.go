package api

import (
	"github.com/gin-gonic/gin"
)

func runV1(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	authGroup(v1)
}

func RunEngine() *gin.Engine {
	engine := gin.Default()
	runV1(engine)
	return engine
}
