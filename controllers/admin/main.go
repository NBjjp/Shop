package admin

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"shop/models"
)

type MainController struct {
	BaseController
}

func (con MainController) Index(ctx *gin.Context) {
	//获取userinfo对应的session
	session := sessions.Default(ctx)
	userinfo := session.Get("userinfo")
	//类型断言 判断userinfo是不是一个string类型
	userinfoStr, ok := userinfo.(string)
	if ok {
		//获取用户信息
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		//2、获取所有的权限
		accessList := []models.Access{}
		models.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort DESC").Find(&accessList)
		//3、获取当前角色拥有的权限 ，并把权限id放在一个map对象里面
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}
		//4、循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
		for i := 0; i < len(accessList); i++ {
			if _, ok := roleAccessMap[accessList[i].Id]; ok {
				accessList[i].Checked = true
			}
			for j := 0; j < len(accessList[i].AccessItem); j++ {
				if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
					accessList[i].AccessItem[j].Checked = true
				}
			}
		}
		ctx.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   userinfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    userinfoStruct[0].IsSuper,
		})
	} else {
		ctx.Redirect(302, "/admin/login")
	}

}

func (con MainController) Welcome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

//公共修改状态的方法
func (con MainController) ChangeStatus(ctx *gin.Context) {
	id, err1 := models.Int(ctx.Query("id"))
	if err1 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入数据错误",
		})
		return
	}
	table := ctx.Query("table")
	field := ctx.Query("field")
	err2 := models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err2 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败 请重试",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}

//公共修改状态的方法
func (con MainController) ChangeNum(ctx *gin.Context) {
	id, err := models.Int(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := ctx.Query("table")
	field := ctx.Query("field")
	num := ctx.Query("num")

	err1 := models.DB.Exec("update "+table+" set "+field+"="+num+" where id=?", id).Error
	if err1 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改数据失败",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "修改成功",
		})
	}

}
