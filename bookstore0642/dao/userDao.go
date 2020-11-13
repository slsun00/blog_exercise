package dao

import (
	"bookstore0642/model"
	"bookstore0642/utils"

)

//CheckUserNameAndPassword 查用户名、密码
func CheckUserNameAndPassword(username string,password string)(*model.User,error){
	sqlStr := `select id ,username,password,email
from users where username=? and password=?`
	row := utils.Db.QueryRow(sqlStr)
	user := &model.User{}
	row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	return user,nil

}

//CheckUserName 查用户名、密码
func CheckUserName(username string)(*model.User,error){



	sqlStr := `select id ,username,password,email
from users where username=? and password=?`
	row := utils.Db.QueryRow(sqlStr)

	// 所以查不到的时候，返回的还是 User ,所以查不到的时候，出现 空的 User
	user := &model.User{}
	 row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)

	//// 查找不到，就返回 nil 这是我自己加的
	//err := row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)

	//if err == sql.ErrNoRows {
	//	return nil,nil
	//}
	return user,nil

}

// SaveUser 保存用户
func SaveUser(username string,password string,email string ) error {
	sqlStr := "insert into users(username,password,email)"
	_ ,err := utils.Db.Exec(sqlStr,username,password,email)
	if err != nil {
		return err
	}

	return nil
}

