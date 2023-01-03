package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
)

type GoodsCateController struct {
	BaseController
}

func (con GoodsCateController) Index(ctx *gin.Context) {
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = 0").Preload("GoodsCateItems").Find(&goodsCateList)
	ctx.HTML(http.StatusOK, "admin/goodsCate/index.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}
func (con GoodsCateController) Add(ctx *gin.Context) {
	//获取顶级分类
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = 0").Find(&goodsCateList)
	ctx.HTML(http.StatusOK, "admin/goodsCate/add.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}
func (con GoodsCateController) DoAdd(ctx *gin.Context) {
	title := ctx.PostForm("title")
	pid, err1 := models.Int(ctx.PostForm("pid"))
	link := ctx.PostForm("link")
	template := ctx.PostForm("template")
	subTitle := ctx.PostForm("sub_title")
	keywords := ctx.PostForm("keywords")
	description := ctx.PostForm("description")
	sort, err2 := models.Int(ctx.PostForm("sort"))
	status, err3 := models.Int(ctx.PostForm("status"))
	if err1 != nil || err3 != nil {
		con.Error(ctx, "传入参数类型不正确", "/goodsCate/add")
		return
	}
	if err2 != nil {
		con.Error(ctx, "排序值必须是整数", "/goodsCate/add")
		return
	}
	cateImgDir, _ := models.UploadImg(ctx, "cate_img")
	goodsCate := models.GoodsCate{
		Title:       title,
		Pid:         pid,
		SubTitle:    subTitle,
		Link:        link,
		Template:    template,
		Keywords:    keywords,
		Description: description,
		CateImg:     cateImgDir,
		Sort:        sort,
		Status:      status,
		AddTime:     int(models.GetUnix()),
	}
	err := models.DB.Create(&goodsCate).Error
	if err != nil {
		con.Error(ctx, "增加数据失败", "/admin/goodsCate/add")
		return
	}
	con.Success(ctx, "增加数据成功", "/admin/goodsCate")
}

func (con GoodsCateController) Edit(ctx *gin.Context) {
	//获取要修改的数据
	id, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "传入参数错误", "/admin/goodsCate")
		return
	}
	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)

	//获取顶级分类
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = 0").Find(&goodsCateList)
	ctx.HTML(http.StatusOK, "admin/goodsCate/edit.html", gin.H{
		"goodsCate":     goodsCate,
		"goodsCateList": goodsCateList,
	})
}
func (con GoodsCateController) DoEdit(ctx *gin.Context) {
	id, err1 := models.Int(ctx.PostForm("id"))
	title := ctx.PostForm("title")
	pid, err2 := models.Int(ctx.PostForm("pid"))
	link := ctx.PostForm("link")
	template := ctx.PostForm("template")
	subTitle := ctx.PostForm("sub_title")
	keywords := ctx.PostForm("keywords")
	description := ctx.PostForm("description")
	sort, err3 := models.Int(ctx.PostForm("sort"))
	status, err4 := models.Int(ctx.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		con.Error(ctx, "传入参数类型不正确", "/goodsCate/add")
		return
	}
	if err3 != nil {
		con.Error(ctx, "排序值必须是整数", "/goodsCate/add")
		return
	}
	cateImgDir, _ := models.UploadImg(ctx, "cate_img")

	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)
	goodsCate.Title = title
	goodsCate.Pid = pid
	goodsCate.Link = link
	goodsCate.Template = template
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	goodsCate.Sort = sort
	goodsCate.Status = status
	if cateImgDir != "" {
		goodsCate.CateImg = cateImgDir
	}
	err := models.DB.Save(&goodsCate).Error
	if err != nil {
		con.Error(ctx, "修改失败", "/admin/goodsCate/edit?id="+models.String(id))
		return
	}
	con.Success(ctx, "修改成功", "/admin/goodsCate")
}
func (con GoodsCateController) Delete(ctx *gin.Context) {
	id, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "传入数据错误", "/admin/goodsCate")
	} else {
		//获取我们要删除的数据
		goodsCate := models.GoodsCate{Id: id}
		models.DB.Find(&goodsCate)
		if goodsCate.Pid == 0 { //顶级分类
			goodsCateList := []models.GoodsCate{}
			models.DB.Where("pid = ?", goodsCate.Id).Find(&goodsCateList)
			if len(goodsCateList) > 0 {
				con.Error(ctx, "当前分类下面子分类，请删除子分类作以后再来删除这个数据", "/admin/goodsCate")
			} else {
				models.DB.Delete(&goodsCate)
				con.Success(ctx, "删除数据成功", "/admin/goodsCate")
			}
		} else { //操作 或者菜单
			models.DB.Delete(&goodsCate)
			con.Success(ctx, "删除数据成功", "/admin/goodsCate")
		}

	}
}
