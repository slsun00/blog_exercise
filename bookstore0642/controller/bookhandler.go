package controller

import (
	"bookstore0642/dao"
	"bookstore0642/model"
	"net/http"
	"html/template"
	"strconv"
)


//GetBooks 获取所有图书
func GetBooks(w http.ResponseWriter,r *http.Request){
	// 获取所有图书
	books,_ := dao.GetBooks()
	// 解析模板文件
	t := template.Must(template.ParseFiles("/views/pages/manager/book_manager.html"))

	// 执行
	t.Execute(w,books)
}

// AddBook 添加图书
func AddBook(w http.ResponseWriter,r *http.Request){
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	//price := r.PostFormValue("price") 字符串
	//sales := r.PostFormValue("sales")
	//stock := r.PostFormValue("stock")


	// 价格、销量、库存
	price , _ := strconv.ParseFloat("price",64)
	sales , _ := strconv.ParseInt("sales",10,64)
	stock , _ := strconv.ParseInt("stock",10,0)


	// 创建 book
	book := &model.Book{
		Title :title,
		Author : author,
		Price : price,
		Sales : int(sales),
		Stock : int(stock),
		ImgPath : "/static/img/default.jpg",
	}

	// 添加 book

	dao.AddBooks(book)

	// 调用GetBooks处理器函数，再次查询一次
	GetBooks(w,r)
}

func DeleteBook(w http.ResponseWriter,r *http.Request) {
	// 获取删除的 id
	bookID := r.FormValue("bookId")
	// 调用删除图书
	dao.DeleteBook(bookID)

	// 调用GetBooks处理器函数，再次查询一次
	GetBooks(w,r)
}

