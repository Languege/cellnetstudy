package main

import (
	"fmt"
	"sync"
)

var (
	// 逻辑中使用的某个变量
	count int

	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
)

func GetCount() int {

	// 锁定
	countGuard.Lock()

	// 在函数退出时解除锁定
	defer countGuard.Unlock()

	return count
}

func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func main() {

	exit := make(chan int)

	// 可以进行并发安全的设置
	go func() {
		for i := 0; i < 10; i++ {
			SetCount(i)
		}

		exit <- 0
	}()

	// 可以进行并发安全的获取
	//

	for i := 0; i < 10; i++ {
		fmt.Println(GetCount())
	}

	<-exit
}
