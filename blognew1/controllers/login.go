package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// 登陆请求
type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get(){
	isExit := this.Input().Get("exit")
	if isExit == "true" {
		// 重新设置 cookie ,删除 cookie ,返回首页
		this.Ctx.SetCookie("userName","",-1,"/")
		this.Ctx.SetCookie("userPwd","",-1,"/")
		this.Redirect("/",301)
		return
	}

	this.TplName = "login.html"
}

// 待随后拆分 ：登陆 、存放密码
func (this *LoginController) Post(){
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	//return


	// 管理员账户登陆
	userName := this.Input().Get("userName")
	userPwd := this.Input().Get("userPwd")
	autoLogin := this.Input().Get("autoLogin")

	if beego.AppConfig.String(" userName") == userName &&
		beego.AppConfig.String("userPwd") == userPwd {
	//	密码未加密保存
		maxAge := 0
		if autoLogin == "on" {
			maxAge = 1 << 31 - 1
		}
		this.Ctx.SetCookie("userName",userName,maxAge,"/")
		this.Ctx.SetCookie("userPwd",userPwd,maxAge,"/")

	}
	this.Redirect("/",301)
	return
}


func CheckAccount (ctx *context.Context) (bool){
	ck , err := ctx.Request.Cookie("userName")
	if err != nil {
		return  false
	}
	userName := ck.Value

	ck , err = ctx.Request.Cookie("userPwd")
	if err != nil {
		return  false
	}

	userPwd := ck.Value
	// cookie 中判断
	return beego.AppConfig.String(" userName") == userName &&
		beego.AppConfig.String("userPwd") == userPwd

}