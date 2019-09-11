package ringbuffer

import (
    "fmt"
    "testing"
    "time"
)

func channelTest(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
            case c <- x:
                x, y = y, x+y
            case <-quit:
                fmt.Println("quit")
                return
        }
    }
}

func TestChannelSelect(t *testing.T) {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
            //比如这里是10个耗时的http请求
            //我们可以同时开启十个协程，最后耗时估计是一个http请求稍微增加一点点(开启通道和调度的时间，几乎可以略微不计)
            //如果是php的话？如何处理(multi)
        }
        quit <- 0
    }()
    channelTest(c, quit)
}

func TestChannelRange(t *testing.T) {
    c := make(chan int)
    go func(in <-chan int) {
        // Using for-range to exit goroutine
        // range has the ability to detect the close/end of a channel
        for x := range in {
            fmt.Printf("Process %d\n", x)
        }
        fmt.Println("over")
    }(c)

    for i := 0; i < 10; i++ {
        c <- i
        if i > 5 {
            close(c)
            break
        }
    }
    //休眠2秒
    time.Sleep(1*time.Second)
    //close(c)
}

//go 常用组合协程配合方案
//go func() do some thing //开启协程
//for select 捕获想要捕获的通道内容进行处理
//for select 是阻塞的
