package dao


import (
	"bookstore0642/model"
	"fmt"
	"testing"
)


func TestBook(t *testing.T) {
	fmt.Println("测试book中的方法")
	t.Run("测试book 的方法",testGetBooks)
}

func testGetBooks(t *testing.T) {
	books ,_ := GetBooks()
	for v,k := range books {
		fmt.Printf("第 %s 本书的信息是 %v \n",v+1,k)
	}
}

func TestAddBooks(t *testing.T) {
	book := &model.Book{
		Title :"三国演义",
		Author : "罗贯中",
		Price : 66.66,
		Sales : 55,
		Stock : 45,
		ImgPath : "/static/img/default.jpg",
	}
	AddBooks(book)

	// 再选择一个函数，查找三国演义这本书
	// getbook(name) 查看是否查找到这本书
}

func TestDeleteBook(t *testing.T) {
	DeleteBook("34")
	// 查询这个id 是否存在
}