/*
	context 官方建议
		1. 不要再任何struct中存放context, 使用函数的第一个参数传递
		2. 不用使用nil的context，如果不知道使用什么context请使用TODO
		3. 不要用来传递函数的可选参数
		4. 相同的context可以传给不同goroutine，它是并发安全的
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	exampleCancel()
	exampleDeadline()
	exampleTimeout()
	exampleValue()
}

func exampleCancel() {
	var wg sync.WaitGroup
	pctx, cancel := context.WithCancel(context.Background())

	wg.Add(4)
	for i := 0; i < 4; i++ {
		ctx, _ := context.WithCancel(pctx)
		go func(ctx context.Context, n int) {
			fmt.Printf("Task %d started.\n", n)
		loop:
			for {
				select {
				case <-time.After(time.Duration(n) * time.Second):
					break loop
				case <-ctx.Done():
					fmt.Printf("Task %d got killed!.\n", n)
					// prevent deadlock
					wg.Done()
					return
				}
			}
			fmt.Printf("Task %d end.\n", n)
			wg.Done()
		}(ctx, 4-i)
	}
	// after 3 sec, cancel it
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	wg.Wait()
}

/*
	output:
		Task 4 started.
		Task 1 started.
		Task 2 started.
		Task 3 started.
		Task 1 end.
		Task 2 end.
		Task 3 end.
		Task 4 got killed!.
		分析:
			cancel context 是基本的context形式相比最原始的空context(Background/TODO)多了cancel方法
			context的用法就是从最开始的context开始往下传，每遇到goroutine就给它分发“继承”的子context
			(当然根据具体业务逻辑去调整，也可以使用同一个context)，当父context调用cancel，所有其下的
			context都会被cancel，子goroutine通过检查Done()返回的通道，能接收就表示被终止了，执行终止
			的相关逻辑。
*/

func exampleDeadline() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	end := make(chan bool)
	go func(ctx context.Context) {
		fmt.Println("I want to sleep for 10 secs")
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("Wonderful sleep")
		case <-ctx.Done():
			fmt.Println("Ooops, your boss wake you up..")
		}
		end <- true
	}(ctx)
	<-end
}

/*
	output:
		I want to sleep for 10 secs
		Ooops, you boss wake you up..
	分析：用具体的时间来触发cancel操作
*/

func exampleTimeout() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	// Timeout context started right after generation
	time.Sleep(3 * time.Second)
	end := make(chan bool)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("I m full of energy!!")
			end <- true
			return
		default:
		}
		fmt.Println("I want to sleep for 10 secs")
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("Wonderful sleep")
		case <-ctx.Done():
			fmt.Println("Ooops, your boss wake you up..")
		}
		end <- true
	}(ctx)
	<-end
}

/*
	output:
		I m full of energy!!
	分析：
		用具体执行的时间来触发cancel操作, 需要注意的是，ctx在创建的时候就会开始算时间
		所以代码上创建了ctx就不要再操作前做耗时操作
*/

type MyKey string

func exampleValue() {
	ctx_parent := context.WithValue(context.Background(), MyKey("name"), "sam")
	fmt.Println("ctx_parent ")
	showValueOfCtx(ctx_parent, MyKey("name"))

	ctx_rewrite := context.WithValue(ctx_parent, MyKey("name"), "jack")
	fmt.Println("ctx_rewrite ")
	showValueOfCtx(ctx_rewrite, MyKey("name"))

	ctx_add := context.WithValue(ctx_parent, MyKey("sex"), "male")
	fmt.Println("ctx_add ")
	showValueOfCtx(ctx_add, MyKey("name"))
	showValueOfCtx(ctx_add, MyKey("sex"))
}
func showValueOfCtx(ctx context.Context, k MyKey) {
	if v := ctx.Value(k); v != nil {
		fmt.Printf("Found value: %s - %s \n", k, v)
		return
	}
	fmt.Printf("Value not found\n")
}

/*
	output:
		ctx_parent
		Found value: name - sam
		ctx_rewrite
		Found value: name - jack
		ctx_add
		Found value: name - sam
		Found value: sex - male
	分析：
		每创建一个context可以携带一个value，后面的创建同名的会把前面的覆盖掉，如果一路“继承”下来
		Value方法会沿着链路往上找返回最近的一个结果
*/
