package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleController struct{}

func (con RoleController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/role/index.html", gin.H{})
}
func (con RoleController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}
func (con RoleController) Edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/role/edit.html", gin.H{})
}
func (con RoleController) DoAdd(ctx *gin.Context) {
	ctx.String(http.StatusOK, "执行添加", gin.H{})
}
func (con RoleController) DoEdit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "执行删除", gin.H{})
}
func (con RoleController) Delete(ctx *gin.Context) {
	ctx.String(http.StatusOK, "执行删除")
}
