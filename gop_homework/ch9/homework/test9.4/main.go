//练习 9.4: 创建一个流水线程序，支持用channel连接任意数量的goroutine，
// 在跑爆内存之前，可以创建多少流水线阶段？一个变量通过整个流水线需要用多久
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	for i :=0;i < 3;i++{
		go func(i int) {
			fmt.Println()
			ch <- i
		}(i)
	}
	fmt.Println("d:",<-ch)
	time.Sleep(1e10)
}
