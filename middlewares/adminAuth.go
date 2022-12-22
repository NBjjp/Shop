package middlewares

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"shop/models"
	"strings"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	//判断用户是否登录   没有登录的用户不能进入后台管理平台

	//获取url访问地址
	pathname := strings.Split(c.Request.URL.String(), "?")[0]

	//获取session中保存的信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//类型断言判断userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)
	if ok {
		//判断userinfo中是否含有账号信息
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		if len(userinfoStruct) == 0 || userinfoStruct[0].Username == "" {
			if pathname != "/admin/login" && pathname != "/admin/captcha" && pathname != "/admin/dologin" {
				c.Redirect(302, "/admin/login")
			}
		}
	} else {
		//没有则跳转到用户登录页面,跳转前判断用户是否在登录页面
		if pathname != "/admin/login" && pathname != "/admin/captcha" && pathname != "/admin/dologin" {
			c.Redirect(302, "/admin/login")
		}
	}
}
