package ringbuffer

import (
    "testing"
    "fmt"
    "time"
)

func TestRingBuffer_Write(t *testing.T) {
    ring := MakeRingBuffer(22)

    go func() {
        for i:=0; i<500;i++ {
                ring.Write(i)
                //fmt.Println(err)
        }
    }()

    //go func() {
    //    fmt.Println("xxx")
    //
    //}()
    for i:=0; i<500;i++ {
        p,v := ring.Read()
        fmt.Println(p, v)
    }
    time.Sleep(2*time.Second)
}