package dao

import (
	"ChatRoom/Go/common/userinfo"
	"errors"
)

var (
	ERROR_USER_NOTEXIST = errors.New("用户不存在")
	ERROR_USER_EXIST    = errors.New("用户已经存在")
	ERROR_USER_PWD      = errors.New("密码不正确")
)

// 服务器启动后初始化一个全局的UserDao
var (
	MyUserDao = &RedisUserDao{}
	MySmsDao  = &MySqlSmsDao{}
)

type IUserDao interface {
	Insert(int, string, string) error        // 添加用户
	Delete(int, string) error                // 删除用户  需要提供ID和密码
	GetUserByID(int) (*userinfo.User, error) // 根据ID获取用户
	IsExist(int) bool                        // 查看用户是否存在

	Update() // 更新用户信息
}

type ISmsDao interface {
	Deposite()
	Withdraw()
}
