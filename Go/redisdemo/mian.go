package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
)

func main() {
	testRedisPool()
}

// 初始化连接池
// 关闭连接池后 获取的conn不能使用
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲数
		MaxActive:   0,   // 最大连接数
		IdleTimeout: 200, // 最大空闲事件
		Dial: func() (redis.Conn, error) { // 创建连接的函数
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

// testRedisPool 测试连接池子
func testRedisPool() {
	// 取出连接
	conn := pool.Get()

	// 延时关闭
	defer conn.Close()

	_, err := conn.Do("Set", "Name", "TomeCate")
	if err != nil {
		fmt.Println("operation failed, err=", err.Error())
		return
	}

	r, err := redis.String(conn.Do("get", "Name"))
	if err != nil {
		fmt.Println("operation failed, err=", err.Error())
		return
	}

	fmt.Println("r=", r)
}

// test  测试操作redis数据
func test() {
	// 连接redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("连接redis失败, err=", err.Error())
		return
	}

	// 延迟断开连接
	defer conn.Close()

	// 通过go向redis写入数据
	_, err = conn.Do("set", "name", "tom")
	if err != nil {
		fmt.Println("操作失败, err = ", err.Error())
	} else {
		fmt.Println("操作成功")
	}

	// 获取redis中的数据
	// 使用redis包中的内置string方法转换结果
	reply, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("操作失败, err = ", err.Error())
	}
	fmt.Println("操作成功, reply=", reply)

	// 通过go向redis写入数据
	_, err = conn.Do("hmset", "user1", "name", "jack", "age", 18)
	if err != nil {
		fmt.Println("操作失败, err = ", err.Error())
	} else {
		fmt.Println("操作成功")
	}

	// 获取redis中的数据
	// 使用redis包中的内置string方法转换结果
	r, err := redis.Strings(conn.Do("hmget", "user1", "name", "age"))
	if err != nil {
		fmt.Println("操作失败, err = ", err.Error())
	}
	fmt.Println("操作成功, r=", r)
}
