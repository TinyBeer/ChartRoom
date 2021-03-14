package main

import (
	"ChartRoom/menuDemo/menu"
)

func main() {

	pMgr := menu.NewPageMgr()
	p := pMgr.AddPage("MainPage", "", "------------欢迎登录海量用户聊天系统------------", "")
	p.AddOption("\t\t登录聊天室", func() {
		pMgr.TurnToPage("HallPage")
	})
	p.AddOption("\t\t注册用户", nil)
	p.AddOption("\t\t退出系统", nil)
	p = pMgr.AddPage("HallPage", "-----聊天室大厅界面-----", "恭喜xxx登录成功", "MainPage")
	p.AddOption("\t在线用户列表", nil)
	p.AddOption("\t发送消息", nil)
	p.AddOption("\t信息列表", nil)
	p.AddOption("\t退出聊天室", nil)

	pMgr.Run()
}
