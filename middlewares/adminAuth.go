package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
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
		} else {
			//用户登录成功 权限判断
			urlPath := strings.Replace(pathname, "/admin/", "", 1)
			if userinfoStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {
				// 1、根据角色获取当前角色的权限列表,然后把权限id放在一个map类型的对象里面
				roleAccessList := []models.RoleAccess{}
				models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccessList)
				roleAccessMap := make(map[int]int)
				for _, v := range roleAccessList {
					roleAccessMap[v.AccessId] = v.AccessId
				}
				// 2、获取当前访问的url对应的权限id 判断权限id是否在角色对应的权限
				access := models.Access{}
				models.DB.Where("url = ?", urlPath).Find(&access)
				//3、判断当前访问的url对应的权限id 是否在权限列表的id中
				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(200, "没有权限")
					c.Abort()
				}
			}
		}
	} else {
		//没有则跳转到用户登录页面,跳转前判断用户是否在登录页面
		if pathname != "/admin/login" && pathname != "/admin/captcha" && pathname != "/admin/dologin" {
			c.Redirect(302, "/admin/login")
		}
	}
}

//排除权限判断的方法
func excludeAuthPath(urlPath string) bool {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()
	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
