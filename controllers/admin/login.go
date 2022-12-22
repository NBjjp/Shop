package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con LoginController) DoLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	captchaId := ctx.PostForm("captchaId")
	VerifyCode := ctx.PostForm("verifyValue")
	//1、验证验证码是否正确
	if flag := models.VerifyCapt(captchaId, VerifyCode); flag {
		//2、如果验证码正确，从数据库中查找username和password，验证是否存在
		userinfo := []models.Manager{}
		password = models.Md5(password)
		models.DB.Where("username = ? and password = ?", username, password).Find(&userinfo)
		//3、如果存在且匹配，将用户信息存入session中，返回登陆成功页面
		if len(userinfo) > 0 {
			session := sessions.Default(ctx)
			//session.Set()不能将结构体切片设置为session，需要将结构体转换成字符串
			userinfoSlice, _ := json.Marshal(userinfo)
			session.Set("userinfo", string(userinfoSlice))
			session.Save()
			con.Success(ctx, "登录成功", "/admin")
		} else {
			//4、登录失败，返回登录页面
			con.Error(ctx, "用户名或密码错误", "/admin/login")
		}

	} else {
		con.Error(ctx, "验证码验证失败", "/admin/login")
	}

}
func (con LoginController) LoginOut(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("userinfo")
	session.Save()
	con.Success(ctx, "您已成功退出", "/admin/loginout")
}

func (con LoginController) Captcha(ctx *gin.Context) {
	id, b64s, err := models.MakeCapt()
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId":  id,
		"captchaImg": b64s,
	})
}
