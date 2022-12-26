package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	"strings"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(ctx *gin.Context) {
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)
	ctx.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}
func (con ManagerController) Add(ctx *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	ctx.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}
func (con ManagerController) DoAdd(ctx *gin.Context) {
	roleId, err1 := models.Int(ctx.PostForm("role_id"))
	if err1 != nil {
		con.Error(ctx, "传入数据错误", "/admin/manager")
		return
	}
	username := strings.Trim(ctx.PostForm("username"), " ")
	password := strings.Trim(ctx.PostForm("password"), " ")
	mobile := strings.Trim(ctx.PostForm("mobile"), " ")
	email := strings.Trim(ctx.PostForm("email"), " ")
	if len(username) < 2 || len(password) < 6 {
		con.Error(ctx, "用户名或密码长度不合法", "/admin/manager/add")
		return
	}
	//判断管理员是否存在
	managerList := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.Error(ctx, "此管理员已经存在", "/admin/manager/add")
		return
	}
	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Mobile:   mobile,
		Email:    email,
		Status:   1,
		RoleId:   roleId,
		AddTime:  int(models.GetUnix()),
	}
	err2 := models.DB.Create(&manager).Error
	if err2 != nil {
		con.Error(ctx, "增加管理员失败", "/admin/manager/add")
	} else {
		con.Success(ctx, "增加管理员成功", "/admin/manager")
	}
}
func (con ManagerController) Edit(ctx *gin.Context) {
	//获取管理员
	id, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx, "数据传入错误", "/admin/manager")
		return
	}
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	//获取所有的角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)

	ctx.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}
func (con ManagerController) DoEdit(ctx *gin.Context) {
	id, err1 := models.Int(ctx.PostForm("id"))
	if err1 != nil {
		con.Error(ctx, "数据传入错误", "/admin/manager")
		return
	}
	roleId, err2 := models.Int(ctx.PostForm("role_id"))
	if err2 != nil {
		con.Error(ctx, "传入数据错误", "/admin/manager")
		return
	}
	username := strings.Trim(ctx.PostForm("username"), " ")
	password := strings.Trim(ctx.PostForm("password"), " ")
	mobile := strings.Trim(ctx.PostForm("mobile"), " ")
	email := strings.Trim(ctx.PostForm("email"), " ")
	//执行修改
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	manager.Username = username
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = roleId
	//判断密码是否为空，为空表示不修改密码，不为空表示修改密码
	if password != "" {
		//判断密码长度是否合法
		if len(password) < 6 {
			con.Error(ctx, "密码长度不合法，密码长度不能小于6位", "/admin/manager/edit?id="+models.String(id))
			return
		}
		manager.Password = models.Md5(password)
	}
	err3 := models.DB.Save(&manager).Error
	if err3 != nil {
		con.Error(ctx, "修改数据失败", "/admin/manager/edit?id="+models.String(id))
		return
	}
	con.Success(ctx, "修改数据成功", "/admin/manager")
}
func (con ManagerController) Delete(ctx *gin.Context) {
	ctx.String(http.StatusOK, "login index")
}
