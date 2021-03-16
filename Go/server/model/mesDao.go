package model

import (
	"log"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

func (udao *UserDao) DepositUserOfflineMesById(id int, data []byte) (err error) {
	// 从连接池取出连接
	conn := udao.Pool.Get()
	// 延时关闭
	defer conn.Close()

	// 将数据存入mesList[userId]中
	res, err := redis.Int64(conn.Do("lpush", "mesList"+strconv.Itoa(id), string(data)))
	log.Println(res)
	if err != nil {
		return err
	}
	// 退出
	return
}

func (udao *UserDao) WithdrawOfflineMesById(id int) (dataSlice []string, err error) {
	// 从连接池取出连接
	conn := udao.Pool.Get()
	// 延时关闭
	defer conn.Close()

	// 将数据存入mesList[userId]中
	dataSlice, err = redis.Strings(conn.Do("lrange", "mesList"+strconv.Itoa(id), 0, -1))
	log.Println(dataSlice)
	if err != nil {
		return
	}

	// 如果留言数量不为零
	if len(dataSlice) != 0 {
		_, err = conn.Do("del", "mesList"+strconv.Itoa(id))
		if err != nil {
			log.Println(err.Error())
		}
	}

	// 退出
	return
}
