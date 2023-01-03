package admin

import (
	"net/http"
	"shop/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type GoodsTypeController struct {
	BaseController
}

func (con GoodsTypeController) Index(ctx *gin.Context) {
	goodsTypeList := []models.GoodsType{}
	models.DB.Find(&goodsTypeList)
	ctx.HTML(http.StatusOK, "admin/goodsType/index.html", gin.H{
		"goodsTypeList": goodsTypeList,
	})

}
func (con GoodsTypeController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/goodsType/add.html", gin.H{})
}

func (con GoodsTypeController) DoAdd(ctx *gin.Context) {

	title := strings.Trim(ctx.PostForm("title"), " ")
	description := strings.Trim(ctx.PostForm("description"), " ")
	status, err1 := models.Int(ctx.PostForm("status"))

	if err1 != nil {
		con.Error(ctx, "传入的参数不正确", "/admin/goodsType/add")
		return
	}

	if title == "" {
		con.Error(ctx, "标题不能为空", "/admin/goodsType/add")
		return
	}
	goodsType := models.GoodsType{
		Title:       title,
		Description: description,
		Status:      status,
		AddTime:     int(models.GetUnix()),
	}

	err := models.DB.Create(&goodsType).Error
	if err != nil {
		con.Error(ctx, "增加商品类型失败 请重试", "/admin/goodsType/add")
	} else {
		con.Success(ctx, "增加商品类型成功", "/admin/goodsType")
	}

}
func (con GoodsTypeController) Edit(ctx *gin.Context) {

	id, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "传入数据错误", "/admin/goodsType")
	} else {
		goodsType := models.GoodsType{Id: id}
		models.DB.Find(&goodsType)
		ctx.HTML(http.StatusOK, "admin/goodsType/edit.html", gin.H{
			"goodsType": goodsType,
		})
	}

}
func (con GoodsTypeController) DoEdit(ctx *gin.Context) {

	id, err1 := models.Int(ctx.PostForm("id"))
	title := strings.Trim(ctx.PostForm("title"), " ")
	description := strings.Trim(ctx.PostForm("description"), " ")
	status, err2 := models.Int(ctx.PostForm("status"))
	if err1 != nil || err2 != nil {
		con.Error(ctx, "传入数据错误", "/admin/goodsType")
		return
	}

	if title == "" {
		con.Error(ctx, "商品类型的标题不能为空", "/admin/goodsType/edit?id="+models.String(id))
	}
	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status

	err3 := models.DB.Save(&goodsType).Error
	if err3 != nil {
		con.Error(ctx, "修改数据失败", "/admin/goodsType/edit?id="+models.String(id))
	} else {
		con.Success(ctx, "修改数据成功", "/admin/goodsType")
	}

}
func (con GoodsTypeController) Delete(ctx *gin.Context) {
	id, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "传入数据错误", "/admin/goodsType")
	} else {
		goodsType := models.GoodsType{Id: id}
		models.DB.Delete(&goodsType)
		con.Success(ctx, "删除数据成功", "/admin/goodsType")
	}
}
