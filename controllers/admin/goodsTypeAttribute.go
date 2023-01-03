package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	"strings"
)

type GoodsTypeAttributeController struct {
	BaseController
}

func (con GoodsTypeAttributeController) Index(ctx *gin.Context) {
	//获取商品类型id
	cateId, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "传入的参数不正确", "/admin/goodsType")
		return
	}
	//获取商品类型属性
	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	models.DB.Where("cate_id = ?", cateId).Find(&goodsTypeAttributeList)
	//获取商品类型属性对应的类型
	goodsType := models.GoodsType{}
	models.DB.Where("id=?", cateId).Find(&goodsType)

	ctx.HTML(http.StatusOK, "admin/goodsTypeAttribute/index.html", gin.H{
		"cateId":                 cateId,
		"goodsTypeAttributeList": goodsTypeAttributeList,
		"goodsType":              goodsType,
	})
}
func (con GoodsTypeAttributeController) Add(ctx *gin.Context) {
	//获取商品类型id
	cateId, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "传入的参数不正确", "/admin/goodsType")
		return
	}
	//获取所有的商品类型
	goodsTypeList := []models.GoodsType{}
	models.DB.Find(&goodsTypeList)
	ctx.HTML(http.StatusOK, "admin/goodsTypeAttribute/add.html", gin.H{
		"goodsTypeList": goodsTypeList,
		"cateId":        cateId,
	})
}

func (con GoodsTypeAttributeController) DoAdd(ctx *gin.Context) {
	title := strings.Trim(ctx.PostForm("title"), " ")
	cateId, err1 := models.Int(ctx.PostForm("cate_id"))
	attrType, err2 := models.Int(ctx.PostForm("attr_type"))
	attrValue := ctx.PostForm("attr_value")
	sort, err3 := models.Int(ctx.PostForm("sort"))

	if err1 != nil || err2 != nil {
		con.Error(ctx, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		con.Error(ctx, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/add?cate_id="+models.String(cateId))
		return
	}

	if err3 != nil {
		con.Error(ctx, "排序值不对", "/admin/goodsTypeAttribute/add?cate_id="+models.String(cateId))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{
		Title:     title,
		CateId:    cateId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		Sort:      sort,
		AddTime:   int(models.GetUnix()),
	}
	err := models.DB.Create(&goodsTypeAttr).Error
	if err != nil {
		con.Error(ctx, "增加商品类型属性失败 请重试", "/admin/goodsTypeAttribute/add?cate_id="+models.String(cateId))
	} else {
		con.Success(ctx, "增加商品类型属性成功", "/admin/goodsTypeAttribute?id="+models.String(cateId))
	}

}
func (con GoodsTypeAttributeController) Edit(ctx *gin.Context) {
	//获取当前要修改数据的id
	id, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "传入数据错误", "/admin/goodsType")
	} else {
		//获取当前id对应的商品类型属性
		goodsTypeAttribute := models.GoodsTypeAttribute{Id: id}
		models.DB.Find(&goodsTypeAttribute)

		//获取所有的商品类型
		goodsTypeList := []models.GoodsType{}
		models.DB.Find(&goodsTypeList)

		ctx.HTML(http.StatusOK, "admin/goodsTypeAttribute/edit.html", gin.H{
			"goodsTypeAttribute": goodsTypeAttribute,
			"goodsTypeList":      goodsTypeList,
		})
	}

}
func (con GoodsTypeAttributeController) DoEdit(ctx *gin.Context) {

	id, err1 := models.Int(ctx.PostForm("id"))
	title := strings.Trim(ctx.PostForm("title"), " ")
	cateId, err2 := models.Int(ctx.PostForm("cate_id"))
	attrType, err3 := models.Int(ctx.PostForm("attr_type"))
	attrValue := ctx.PostForm("attr_value")
	sort, err4 := models.Int(ctx.PostForm("sort"))

	if err1 != nil || err2 != nil || err3 != nil {
		con.Error(ctx, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		con.Error(ctx, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/edit?id="+models.String(id))
		return
	}
	if err4 != nil {
		con.Error(ctx, "排序值不对", "/admin/goodsTypeAttribute/edit?id="+models.String(id))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
	models.DB.Find(&goodsTypeAttr)
	goodsTypeAttr.Title = title
	goodsTypeAttr.CateId = cateId
	goodsTypeAttr.AttrType = attrType
	goodsTypeAttr.AttrValue = attrValue
	goodsTypeAttr.Sort = sort
	err := models.DB.Save(&goodsTypeAttr).Error
	if err != nil {
		con.Error(ctx, "修改数据失败", "/admin/goodsTypeAttribute/edit?id="+models.String(id))
		return
	}
	con.Success(ctx, "需改数据成功", "/admin/goodsTypeAttribute?id="+models.String(cateId))

}
func (con GoodsTypeAttributeController) Delete(ctx *gin.Context) {
	id, err1 := models.Int(ctx.Query("id"))
	cateId, err2 := models.Int(ctx.Query("cate_id"))
	if err1 != nil || err2 != nil {
		con.Error(ctx, "传入参数错误", "/admin/goodsType")
	} else {
		goodsTypeAttr := models.GoodsTypeAttribute{Id: id}
		models.DB.Delete(&goodsTypeAttr)
		con.Success(ctx, "删除数据成功", "/admin/goodsTypeAttribute?id="+models.String(cateId))
	}
}
