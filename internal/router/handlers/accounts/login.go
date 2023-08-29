package accounts

import "github.com/gin-gonic/gin"

func Login(ctx *gin.Context) {

	ctx.String(200, "login")
}
