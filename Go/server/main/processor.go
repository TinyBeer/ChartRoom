package main

import (
	"ChartRoom/common/message"
	"ChartRoom/common/utils"
	"ChartRoom/server/processes"
	"errors"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (pro *Processor) serverProcess(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 创建UserProcess实例
		up := &processes.UserProcess{Conn: pro.Conn}
		up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &processes.UserProcess{Conn: pro.Conn}
		// 处理注册消息
		up.ServerProccessRegister(mes)
	case message.SmsMesType:
		smsProcess := &processes.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	case message.LogoutMesType:
		up := &processes.UserProcess{Conn: pro.Conn}
		up.ServerProcessLogout(mes)
		return errors.New("用户登出")
	default:
		err = errors.New("未知消息类型")
	}

	return
}

func (pro *Processor) Process2() (err error) {
	// 使用Transfer读写数据
	tf := utils.NewTransfer(pro.Conn)
	// 读取客户发送的消息
	for {
		mes, err := tf.ReadPkg()
		if err != nil {

			switch err {
			case io.EOF:
				fmt.Println("客户端断开连接")
			default:
				fmt.Println("客户端连接中断")
			}
			return err
		}

		err = pro.serverProcess(&mes)
		if err != nil {
			fmt.Println("通讯异常")
			// fmt.Println("通讯协程断开， err=", err.Error())
		}
	}
}
