package adapters

import "github.com/gin-gonic/gin"

func Adapter(handler Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(
			HTTPContext{
				ResponseWriter: ctx.Writer,
				Request:        ctx.Request,
			},
		)
	}
}
