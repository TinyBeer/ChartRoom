package model

import (
	"ChartRoom/common/message"
	"net"
)

// 创建全局
type CurUser struct {
	message.User
	Conn net.Conn
}
