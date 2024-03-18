package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "return for ctx.Done()")
			return
		default:
			fmt.Println(name, "working")
			time.Sleep(1 * time.Second)
		}
	}
}

func manager(cancel func()) {
	time.Sleep(882 * time.Second)
	fmt.Println("manager called cancel()")
	cancel()
}

func main() {
	const MAX_WORKING_DURATION = 5 * time.Second

	ctxWithCancel, cancel := context.WithTimeout(context.Background(), MAX_WORKING_DURATION)
	go worker(ctxWithCancel, "worker1")
	go worker(ctxWithCancel, "worker2")
	go manager(cancel)
	<-ctxWithCancel.Done() // 暂停1秒便于协程的打印输出
	time.Sleep(1 * time.Second)
	fmt.Println("company closed")
}
