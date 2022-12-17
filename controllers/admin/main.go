package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainController struct {
	BaseController
}

func (con MainController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}

func (con MainController) Welcome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
