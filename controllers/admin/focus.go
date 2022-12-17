package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FocusController struct{}

func (con FocusController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/index.html", gin.H{})
}
func (con FocusController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}
func (con FocusController) Edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{})
}
func (con FocusController) Delete(ctx *gin.Context) {
	ctx.String(http.StatusOK, "login index")
}
