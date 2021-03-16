package processes

import (
	"ChartRoom/common/message"
	"ChartRoom/common/utils"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

// 声明结构体
type SmsProcess struct {
	//..
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
