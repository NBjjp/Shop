package routers

import (
	"github.com/gin-gonic/gin"
	admincon "shop/controllers/admin"
)

func AdminRoutersInit(r *gin.Engine) {
	//middlewares.InitMiddleware中间件
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", admincon.MainController{}.Index)
		adminRouters.GET("/welcome", admincon.MainController{}.Welcome)

		adminRouters.GET("/login", admincon.LoginController{}.Login)
		adminRouters.GET("/dologin", admincon.LoginController{}.DoLogin)

		adminRouters.GET("/manager", admincon.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admincon.ManagerController{}.Add)
		adminRouters.GET("/manager/edit", admincon.ManagerController{}.Edit)
		adminRouters.GET("/manager/delete", admincon.ManagerController{}.Delete)

		adminRouters.GET("/focus", admincon.ManagerController{}.Index)
		adminRouters.GET("/focus/add", admincon.ManagerController{}.Add)
		adminRouters.GET("/focus/edit", admincon.ManagerController{}.Edit)
		adminRouters.GET("/focus/delete", admincon.ManagerController{}.Delete)

	}
}
