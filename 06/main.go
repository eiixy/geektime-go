package main

import (
	"fmt"
	"time"
)

var (
	window int64 = 3  // 窗口时间长度
	limit  int   = 30 // 允许通过的数量
)

var queue map[string][]int64

func Fun1(name string) bool {
	now := time.Now().Unix()
	if queue == nil {
		queue = make(map[string][]int64)
		queue[name] = append(queue[name], 0)
	}
	fmt.Printf("now :%d ", now)

	// 切片长度不够直接返回
	if len(queue[name]) < limit {
		queue[name] = append(queue[name], now)
		return true
	}

	// 判断当前时间与开始时间的差值是否大于窗口时间长度
	start := queue[name][0]
	if now-start <= window {
		return false
	}

	// 移除最小时间，填充当前时间
	queue[name] = queue[name][1:]
	queue[name] = append(queue[name], now)
	return true

}

func main() {
	for i := 0; i < 100; i++ {
		result := Fun1("test")
		time.Sleep(40 * time.Millisecond)
		fmt.Printf("result: %t\r\n", result)
	}
}
