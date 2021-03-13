package processes

import (
	"ChartRoom/common/message"
	"encoding/json"
	"fmt"
)

// 传入smsMes类型数据
func outputGroupMes(mes *message.Message) {
	// 反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal failed, err=", err.Error())
		return
	}

	fmt.Printf("收到来自用户%d的群发消息:\n", smsMes.UserID)
	fmt.Println(smsMes.Content)
}
