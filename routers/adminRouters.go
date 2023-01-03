package routers

import (
	"github.com/gin-gonic/gin"
	admincon "shop/controllers/admin"
	"shop/middlewares"
)

func AdminRoutersInit(r *gin.Engine) {
	//middlewares.InitMiddleware中间件
	adminRouters := r.Group("/admin", middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/", admincon.MainController{}.Index)
		adminRouters.GET("/welcome", admincon.MainController{}.Welcome)
		adminRouters.GET("/changeStatus", admincon.MainController{}.ChangeStatus)
		adminRouters.GET("/changeNum", admincon.MainController{}.ChangeNum)

		adminRouters.GET("/login", admincon.LoginController{}.Login)
		adminRouters.POST("/dologin", admincon.LoginController{}.DoLogin)
		adminRouters.GET("/loginOut", admincon.LoginController{}.LoginOut)
		adminRouters.GET("/captcha", admincon.LoginController{}.Captcha)

		adminRouters.GET("/manager", admincon.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admincon.ManagerController{}.Add)
		adminRouters.POST("/manager/doAdd", admincon.ManagerController{}.DoAdd)
		adminRouters.GET("/manager/edit", admincon.ManagerController{}.Edit)
		adminRouters.POST("/manager/doEdit", admincon.ManagerController{}.DoEdit)
		adminRouters.GET("/manager/delete", admincon.ManagerController{}.Delete)

		adminRouters.GET("/focus", admincon.FocusController{}.Index)
		adminRouters.GET("/focus/add", admincon.FocusController{}.Add)
		adminRouters.POST("/focus/doAdd", admincon.FocusController{}.DoAdd)
		adminRouters.GET("/focus/edit", admincon.FocusController{}.Edit)
		adminRouters.POST("/focus/doEdit", admincon.FocusController{}.DoEdit)
		adminRouters.GET("/focus/delete", admincon.FocusController{}.Delete)

		adminRouters.GET("/role", admincon.RoleController{}.Index)
		adminRouters.GET("/role/add", admincon.RoleController{}.Add)
		adminRouters.GET("/role/edit", admincon.RoleController{}.Edit)
		adminRouters.POST("/role/doAdd", admincon.RoleController{}.DoAdd)
		adminRouters.POST("/role/doEdit", admincon.RoleController{}.DoEdit)
		adminRouters.GET("/role/delete", admincon.RoleController{}.Delete)
		adminRouters.GET("/role/auth", admincon.RoleController{}.Auth)
		adminRouters.POST("/role/doAuth", admincon.RoleController{}.DoAuth)

		adminRouters.GET("/access", admincon.AccessController{}.Index)
		adminRouters.GET("/access/add", admincon.AccessController{}.Add)
		adminRouters.POST("/access/doAdd", admincon.AccessController{}.DoAdd)
		adminRouters.GET("/access/delete", admincon.AccessController{}.Delete)
		adminRouters.GET("/access/edit", admincon.AccessController{}.Edit)
		adminRouters.POST("/access/doEdit", admincon.AccessController{}.DoEdit)

		adminRouters.GET("/goodsCate", admincon.GoodsCateController{}.Index)
		adminRouters.GET("/goodsCate/add", admincon.GoodsCateController{}.Add)
		adminRouters.POST("/goodsCate/doAdd", admincon.GoodsCateController{}.DoAdd)
		adminRouters.GET("/goodsCate/edit", admincon.GoodsCateController{}.Edit)
		adminRouters.POST("/goodsCate/doEdit", admincon.GoodsCateController{}.DoEdit)
		adminRouters.GET("/goodsCate/delete", admincon.GoodsCateController{}.Delete)

		adminRouters.GET("/goodsType", admincon.GoodsTypeController{}.Index)
		adminRouters.GET("/goodsType/add", admincon.GoodsTypeController{}.Add)
		adminRouters.POST("/goodsType/doAdd", admincon.GoodsTypeController{}.DoAdd)
		adminRouters.GET("/goodsType/edit", admincon.GoodsTypeController{}.Edit)
		adminRouters.POST("/goodsType/doEdit", admincon.GoodsTypeController{}.DoEdit)
		adminRouters.GET("/goodsType/delete", admincon.GoodsTypeController{}.Delete)

		adminRouters.GET("/goodsTypeAttribute", admincon.GoodsTypeAttributeController{}.Index)
		adminRouters.GET("/goodsTypeAttribute/add", admincon.GoodsTypeAttributeController{}.Add)
		adminRouters.POST("/goodsTypeAttribute/doAdd", admincon.GoodsTypeAttributeController{}.DoAdd)
		adminRouters.GET("/goodsTypeAttribute/edit", admincon.GoodsTypeAttributeController{}.Edit)
		adminRouters.POST("/goodsTypeAttribute/doEdit", admincon.GoodsTypeAttributeController{}.DoEdit)
		adminRouters.GET("/goodsTypeAttribute/delete", admincon.GoodsTypeAttributeController{}.Delete)
	}
}
