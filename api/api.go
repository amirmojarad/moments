package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// getUsernameFromCtx get username that pushed in gin ctx in middleware layer
func getUsernameFromCtx(ctx *gin.Context) string {
	return fmt.Sprint(ctx.MustGet("username"))
}

func runV1(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	authGroup(v1)
}

func RunEngine() *gin.Engine {
	engine := gin.Default()
	runV1(engine)
	return engine
}
