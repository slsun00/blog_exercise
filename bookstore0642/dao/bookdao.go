package dao

import (
	"bookstore0642/model"
	"bookstore0642/utils"
)
//GetBooks 获取数据库中的所有图书
func GetBooks() ([]*model.Book ,error){
	sqlStr := `select id,title,author,price,sales,stock,img_path
from books`
	rows,err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil ,err
	}

	var books []*model.Book
	for rows.Next() {
		var book *model.Book

		// 字段赋值
		rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)

		books = append(books,book)
	}

	return books,nil

}


//
func AddBooks(b *model.Book) (err error){
	sqlStr := `insert into books(title,author,price,sales,stock,img_path)
values(?,?,?,?,?)`
	_ , err  =  utils.Db.Exec(sqlStr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.ImgPath)
	if err != nil {
		return err
	}
	return nil

}


func DeleteBook(bookId string) error {
	sqlStr := "delete form books where id = ?"

	_ , err := utils.Db.Exec(sqlStr,bookId)
	if err != nil {
		return nil
	}
	return nil
}