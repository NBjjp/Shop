package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	"strings"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(ctx *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	fmt.Println(roleList)
	ctx.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}
func (con RoleController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}
func (con RoleController) Edit(ctx *gin.Context) {
	id := ctx.Query("id")
	strId, err := models.Int(id)
	if err != nil {
		con.Error(ctx, "id获取失败", "/admin/role")
	} else {
		role := models.Role{Id: strId}
		models.DB.Find(&role)
		ctx.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}

}
func (con RoleController) DoAdd(ctx *gin.Context) {
	title := strings.Trim(ctx.PostForm("title"), " ")
	description := strings.Trim(ctx.PostForm("description"), " ")
	if title == "" {
		con.Error(ctx, "角色标题不能为空，请重新输入", "/admin/role/add")
	} else {
		role := models.Role{}
		role.Title = title
		role.Description = description
		role.Status = 1
		role.Addtime = int(models.GetUnix())
		err := models.DB.Create(&role).Error
		if err != nil {
			con.Error(ctx, "增加角色失败，请重试", "/admin/role/add")
		} else {
			con.Success(ctx, "增加角色成功", "/admin/role")
		}
	}

}
func (con RoleController) DoEdit(ctx *gin.Context) {
	id, err1 := models.Int(ctx.PostForm("id"))
	if err1 != nil {
		con.Error(ctx, "传入数据错误", "/admin/role")
	} else {
		title := strings.Trim(ctx.PostForm("title"), " ")
		description := strings.Trim(ctx.PostForm("description"), " ")
		if title == "" {
			con.Error(ctx, "角色的标题不能为空", "/admin/role/edit")
		} else {
			role := models.Role{Id: id}
			models.DB.Find(&role)
			role.Title = title
			role.Description = description
			err2 := models.DB.Save(&role).Error
			if err2 != nil {
				con.Error(ctx, "修改数据失败", "/admin/role/edit?id="+models.String(id))
			} else {
				con.Success(ctx, "修改数据成功", "/admin/role/edit?id="+models.String(id))
			}
		}
	}

}
func (con RoleController) Delete(ctx *gin.Context) {
	id, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "传入数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Delete(&role)
		con.Success(ctx, "删除成功", "/admin/role")
	}
}
func (con RoleController) Auth(ctx *gin.Context) {
	//获取角色id
	role_id, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "传入数据错误", "/admin/role")
		return
	}
	//获取所有的权限
	accessList := []models.Access{}
	err1 := models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList).Error
	if err1 != nil {
		con.Error(ctx, "获取顶级权限失败", "/admin/access")
		return
	}
	//3、获取当前角色拥有的权限 ，并把权限id放在一个map对象里面
	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id=?", role_id).Find(&roleAccess)
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

	ctx.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"role_id":    role_id,
		"accessList": accessList,
	})

}
func (con RoleController) DoAuth(ctx *gin.Context) {
	//获取角色id
	roleId, err := models.Int(ctx.PostForm("role_id"))
	if err != nil {
		con.Error(ctx, "传入数据错误", "/admin/role")
		return
	}
	//获取权限id
	accessIds := ctx.PostFormArray("access_node[]")
	fmt.Println(accessIds)

	//删除当前角色的权限
	roleAccess := models.RoleAccess{}
	models.DB.Where("role_id=?", roleId).Delete(&roleAccess)
	//增加当前角色对应的权限
	for _, v := range accessIds {
		roleAccess.RoleId = roleId
		accessId, _ := models.Int(v)
		roleAccess.AccessId = accessId
		models.DB.Create(&roleAccess)
	}
	con.Success(ctx, "授权成功", "/admin/role/auth?id="+models.String(roleId))
}
