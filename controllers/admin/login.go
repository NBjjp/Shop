package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con LoginController) DoLogin(ctx *gin.Context) {
	ctx.String(http.StatusOK, "login index")
}
