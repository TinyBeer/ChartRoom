package dao

import (
	"ChatRoom/Go/server/database"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	database.Init("root", "123456", "127.0.0.1", 3306, "chart_room")

	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	database.Close()

	os.Exit(retCode) // 退出测试
}

func TestDepositeByID(t *testing.T) {
	msd := MySqlSmsDao{}
	err := msd.DepositeByID(100, "hello")
	if err != nil {
		t.Error(err)
	}
}

func TestWithdrawByID(t *testing.T) {
	msd := MySqlSmsDao{}
	smsList, err := msd.WithdrawByID(100)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v\n", smsList)
}
