package route

import (
	"bit-particle-cannon/handler"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.POST("/lark", handler.HandleLark)
}
