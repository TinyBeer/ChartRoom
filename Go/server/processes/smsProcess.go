package processes

import (
	"ChartRoom/common/message"
	"ChartRoom/common/utils"
	"ChartRoom/server/model"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

// 声明结构体
type SmsProcess struct {
	//..
}

func (sp *SmsProcess) SendMessage(mes *message.Message) (err error) {

	// 1.取出mes.Data,并反序列化
	var messageMes message.MessageMes
	err = utils.Unpack(mes, &messageMes)
	if err != nil {
		log.Println("ServerProcessMessage utils.Unpack failed, err=", err.Error())
		return
	}
	// 2.判断用户是否在线
	up, ok := userMgr.onlineUsers[messageMes.ToUserID]
	if ok {
		// 4.1 在线  转发消息
		// 构建mes
		var sendMes message.Message
		sendMes.Type = message.SmsMesType

		var smsMes message.SmsMes
		smsMes.Content = messageMes.Content
		smsMes.User = messageMes.User
		// 封包
		err = utils.Pack(&sendMes, &smsMes)
		if err != nil {
			log.Println("ServerProcessMessage utils.Pack failed, err=", err.Error())
			return
		}
		// 使用SendMesToEachOnlineUser函数发送sendMes
		sp.SendMesToEachOnlineUser(&sendMes, up.Conn)
	} else {
		// 4.2 不在线 转存消息
		err = model.MyUserDao.DepositUserOfflineMesById(messageMes.ToUserID, []byte(mes.Data))
		if err != nil {
			log.Println("DepositUserOfflineMesById failed, err=", err.Error())
			return
		}
	}
	return
}

// 转发消息
func (sp *SmsProcess) SendGroupMes(mes *message.Message) (err error) {

	// 取出smsMes
	var smsMes message.SmsMes
	err = utils.Unpack(mes, &smsMes)
	if err != nil {
		log.Println("Unpack failed, err=", err.Error())
		return
	}

	// 遍历服务端的onlineUsers
	// 转发消息
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserID {
			continue
		}
		sp.SendMesToEachOnlineUser(mes, up.Conn)
	}
	return
}

// 发送消息
func (sp *SmsProcess) SendMesToEachOnlineUser(mes *message.Message, conn net.Conn) (err error) {
	tf := utils.NewTransfer(conn)

	// 序列化
	data, err := json.Marshal(&mes)
	if err != nil {
		return
	}

	err = tf.WriteData(data)
	if err != nil {
		fmt.Println("tf.WriteData failed, err=", err.Error())
		return
	}
	return
}
