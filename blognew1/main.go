package main

import (
	_ "blognew1/routers"
	"github.com/astaxie/beego"
)

/*
func init() {
	软件程序打开，就会连接数据库

}
*/


func main() {
	beego.Run()
}

