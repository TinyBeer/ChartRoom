package processes

import (
	"ChartRoom/common/message"
	"ChartRoom/common/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMes(content string) (err error) {
	// 创建一个mes
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content

	smsMes.UserID = CurUser.UserID
	smsMes.UserStatus = CurUser.UserStatus

	// 序列化smsMes
	data, err := json.Marshal(&smsMes)
	if err != nil {
		fmt.Println("json.Marshal failed, err=", err.Error())
		return
	}

	// 装在mes内容
	mes.Data = string(data)

	// 发送
	tf := utils.NewTransfer(CurUser.Conn)
	err = tf.WritePkg(&mes)
	if err != nil {
		fmt.Println("tf.WritePkg failed, err=", err.Error())
		return
	}
	return
}
