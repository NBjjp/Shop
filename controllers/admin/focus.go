package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(ctx *gin.Context) {
	focusList := []models.Focus{}
	models.DB.Find(&focusList)
	ctx.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}
func (con FocusController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}
func (con FocusController) DoAdd(ctx *gin.Context) {
	title := ctx.PostForm("title")
	focusType, err1 := models.Int(ctx.PostForm("focus_type"))
	link := ctx.PostForm("link")
	sort, err2 := models.Int(ctx.PostForm("sort"))
	status, err3 := models.Int(ctx.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(ctx, "非法请求", "/admin/focus/add")
		return
	}
	if err2 != nil {
		con.Error(ctx, "请输入正确的排序值", "/admin/focus/add")
		return
	}
	//上传文件
	focusImgSrc, err4 := models.UploadImg(ctx, "focus_img")
	if err4 != nil {
		con.Error(ctx, err4.Error(), "/admin/focus/add")
		return
	}
	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}
	err5 := models.DB.Create(&focus).Error
	if err5 != nil {
		con.Error(ctx, "增加轮播图失败", "/admin/focus/add")
		return
	} else {
		con.Success(ctx, "增加轮播图成功", "/admin/focus")
	}
}
func (con FocusController) Edit(ctx *gin.Context) {
	id, err1 := models.Int(ctx.Query("id"))
	if err1 != nil {
		con.Error(ctx, "传入参数错误", "/admin/focus")
		return
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	ctx.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}
func (con FocusController) DoEdit(ctx *gin.Context) {
	id, err1 := models.Int(ctx.PostForm("id"))
	title := ctx.PostForm("title")
	focusType, err2 := models.Int(ctx.PostForm("focus_type"))
	link := ctx.PostForm("link")
	sort, err3 := models.Int(ctx.PostForm("sort"))
	status, err4 := models.Int(ctx.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		con.Error(ctx, "非法请求", "/admin/focus")
		return
	}
	if err3 != nil {
		con.Error(ctx, "请输入正确的排序值", "/admin/focus/edit?id="+models.String(id))
		return
	}
	//上传文件
	focusImg, _ := models.UploadImg(ctx, "focus_img")

	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImg != "" {
		focus.FocusImg = focusImg
	}
	err5 := models.DB.Save(&focus).Error
	if err5 != nil {
		con.Error(ctx, "修改数据失败请重新尝试", "/admin/focus/edit?id="+models.String(id))
		return
	} else {
		con.Success(ctx, "增加轮播图成功", "/admin/focus")
	}
}

func (con FocusController) Delete(ctx *gin.Context) {
	id, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "传入数据错误", "/admin/focus")
		return
	} else {
		focus := models.Focus{Id: id}
		models.DB.Delete(&focus)
		//根据自己的需要 要不要删除图片
		// os.Remove("static/upload/20210915/1631694117.jpg")
		con.Success(ctx, "删除数据成功", "/admin/focus")
	}
}
