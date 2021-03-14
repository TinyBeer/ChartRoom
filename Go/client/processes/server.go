package processes

import (
	"ChartRoom/common/message"
	"ChartRoom/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

// 显示登录成功后的界面
func ShowMenu(userName string) {

	fmt.Printf("----恭喜%3s登录成功-----\n", userName)
	fmt.Println("----1：在线用户列表-----")
	fmt.Println("----2：发送消息---------")
	fmt.Println("----3：信息列表---------")
	fmt.Println("----4：退出系统---------")
	fmt.Println("请选择1-4：")

	var key int
	var content string

	// 总会使用SmsProcess
	smsProcess := &SmsProcess{}
	for {
		fmt.Scanln(&key)
		switch key {
		case 1:
			// fmt.Println("显示在线用户列表")
			OutputOnlineUsers()
		case 2:
			fmt.Println("请输入要发送的消息:")
			fmt.Scanf("%s\n", &content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Print("退出系统")
			up := &UserProcess{}
			up.Logout()
			return
		default:
			fmt.Println("输入有误，从新输入：")

		}
	}

}

// 和服务器保持通信
func ServerMesProcess(conn net.Conn) {
	defer conn.Close()
	// 创建一个Transfer 不停的读取消息
	tf := utils.NewTransfer(conn)
	for {
		// fmt.Println("客户端正在读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg failed, err=", err.Error())
			return
		}

		switch mes.Type {
		case message.NotifyUserStatusMesType:
			// 处理用户状态更新消息
			// 取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			// 反序列化
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("json.Unmarshall failed, err=", err.Error())
				continue
			}
			updateUserStatus(&notifyUserStatusMes)
			OutputOnlineUsers()
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("获取到未知消息类型")
		}

	}
}
