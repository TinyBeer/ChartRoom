package main

import (
	"ChartRoom/client/processes"
	"ChartRoom/client/view"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	key      int
	userID   int
	userPwd  string
	userName string
	loop     bool
	cChan    chan net.Conn
)

func initView() *view.PageMgr {
	var content string
	smsProcess := &processes.SmsProcess{}
	pMgr := view.NewPageMgr()
	p := pMgr.AddPage("MainPage", "", "------------欢迎登录海量用户聊天系统------------", "")
	p.AddOption("\t\t登录聊天室", func() {
		fmt.Println("请输入用户ID：")
		fmt.Scanln(&userID)
		fmt.Println("请输入密码：")
		fmt.Scanln(&userPwd)
		up := &processes.UserProcess{}
		conn, err := up.Login(userID, userPwd)
		if err != nil {
			log.Println(err.Error())
			return
		} else {
			// 启一个协程保持和服务器的连接
			cChan <- conn
			pMgr.TurnToPage("HallPage")
		}
	})

	p.AddOption("\t\t注册用户", func() {
		fmt.Println("请输入用户ID：")
		fmt.Scanln(&userID)
		fmt.Println("请输入密码：")
		fmt.Scanln(&userPwd)
		fmt.Println("请输入用户昵称：")
		fmt.Scanln(&userName)
		// 调用UserDao实例  实现注册
		up := &processes.UserProcess{}
		up.Register(userID, userPwd, userName)
	})

	p.AddOption("\t\t退出系统", func() {
		os.Exit(0)
	})

	p = pMgr.AddPage("HallPage", "-----聊天室大厅界面-----", "恭喜xxx登录成功", "MainPage")

	p.AddOption("\t在线用户列表", func() {
		processes.OutputOnlineUsers()
	})
	p.AddOption("\t发送消息", func() {
		fmt.Println("请输入要发送的消息:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	})
	p.AddOption("\t信息列表", func() {

	})
	p.AddOption("\t退出聊天室", func() {
		up := &processes.UserProcess{}
		up.Logout()
		pMgr.GoBack()
	})
	return pMgr
}

func main() {
	cChan = make(chan net.Conn, 1)

	go func() {
		for {
			conn := <-cChan
			go processes.ServerMesProcess(conn)
		}
	}()

	initView().Run()
}
