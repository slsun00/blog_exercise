package controller

import (
	"bookstore0642/dao"
	"net/http"
	"html/template"
)

func Login(w http.ResponseWriter,r *http.Request){
	// 获取用户账号密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user , _ := dao.CheckUserNameAndPassword(username,password)



	// 使用改过的 就可以使用 user == nil
	if user.ID != 0 {
		// 用户名密码正确,转到注册成功页面
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w,"")

	} else {
		// 用户名密码不正确,转到注册成功页面
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w,"用户名或密码不正确!")
	}
}

func Regist (w http.ResponseWriter,r *http.Request) {
	// 获取用户账号密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	// 验证
	user , _ := dao.CheckUserName(username)

	// 使用改过的 就可以使用 user == nil
	if user.ID > 0 {
		// 用户名存在，重新到注册页面
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w,"")

	} else {
		// 用户名不存在
		dao.SaveUser(username,password,email)
		// 用户名存在，重新到注册页面
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w,"")
	}

}

// 通过 ajax 信息
func CheckUserName(w http.ResponseWriter,r *http.Request){
	// 获取用户名
	username := r.PostFormValue("username")
	// 验证
	user , _ := dao.CheckUserName(username)

	// 使用改过的 就可以使用 user == nil
	if user.ID > 0 {
		// 用户名存在
		w.Write([]byte("用户名已存在"))
	} else {
		// 用户名不存在
		w.Write([]byte("<font style='color:green'>用户名可用"))

	}

}

