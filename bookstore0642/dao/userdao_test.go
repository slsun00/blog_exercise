package dao

import (
	"fmt"
	"testing"
)





func TestUser(t *testing.T) {
	t.Run("验证用户名或密码",testLogin)
	t.Run("验证用户",testRegist)
	t.Run("保存用户",testSaveUser)

}

// 子测试

func testLogin (t *testing.T) {
	user,_ := CheckUserNameAndPassword("admin","1234")
	fmt.Println("获取用户信息是：",user)
}

func testRegist (t *testing.T) {
	user,_ := CheckUserName("admin","1234")
	fmt.Println("获取用户信息是：",user)
}

func testSaveUser (t *testing.T) {
	SaveUser("admin","1234","qq.com")
}