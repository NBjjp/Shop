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

		adminRouters.GET("/login", admincon.LoginController{}.Login)
		adminRouters.POST("/dologin", admincon.LoginController{}.DoLogin)
		adminRouters.GET("/loginout", admincon.LoginController{}.LoginOut)
		adminRouters.GET("/captcha", admincon.LoginController{}.Captcha)

		adminRouters.GET("/manager", admincon.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admincon.ManagerController{}.Add)
		adminRouters.GET("/manager/edit", admincon.ManagerController{}.Edit)
		adminRouters.GET("/manager/delete", admincon.ManagerController{}.Delete)

		adminRouters.GET("/focus", admincon.ManagerController{}.Index)
		adminRouters.GET("/focus/add", admincon.ManagerController{}.Add)
		adminRouters.GET("/focus/edit", admincon.ManagerController{}.Edit)
		adminRouters.GET("/focus/delete", admincon.ManagerController{}.Delete)

		adminRouters.GET("/role", admincon.RoleController{}.Index)
		adminRouters.GET("/role/add", admincon.RoleController{}.Add)
		adminRouters.GET("/role/edit", admincon.RoleController{}.Edit)
		adminRouters.POST("/role/doAdd", admincon.RoleController{}.DoAdd)
		adminRouters.POST("/role/doEdit", admincon.RoleController{}.DoEdit)
		adminRouters.GET("/role/delete", admincon.RoleController{}.Delete)

	}
}
