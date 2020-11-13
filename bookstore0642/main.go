package main

import (
	"bookstore0642/controller"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter,r *http.Request) {

	// 解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w,"")
}

func main(){

	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/",http.StripPrefix("/pages/",http.FileServer(http.Dir("views/pages"))))


	http.HandleFunc("/login",controller.Login)

	// 注册
	http.HandleFunc("/regist",controller.Regist)

	// ajax 验证用户名是否可用
	http.HandleFunc("/checkUserName",controller.CheckUserName)

	// 获取图书
	http.HandleFunc("/getBooks",controller.GetBooks)

	// 添加图书

	http.HandleFunc("/addBook",controller.AddBook)
	http.HandleFunc("/deleteBook",controller.DeleteBook)


	// 思考的时候，从上往下思考，写代码从小到上
	http.HandleFunc("/",IndexHandler)
	http.ListenAndServe(":8080",nil)

}