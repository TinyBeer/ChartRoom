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
	err = utils.Pack(&mes, &smsMes)
	if err != nil {
		fmt.Println("Pack failed, err=", err.Error())
		return
	}

	// 序列化
	data, err := json.Marshal(&mes)
	if err != nil {
		return
	}

	// 发送
	tf := utils.NewTransfer(CurUser.Conn)
	err = tf.WriteData(data)
	if err != nil {
		fmt.Println("tf.WriteData failed, err=", err.Error())
		return
	}
	return
}
