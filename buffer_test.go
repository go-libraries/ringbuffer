package ringbuffer

import (
    "testing"
    "fmt"
)

func TestRingBuffer_Write(t *testing.T) {
    ring := MakeRingBuffer(22)

    for i:=0; i<100;i++ {
        go func() {
            ring.Write(i)
            //fmt.Println(e)
        }()
    }

    for i:=0; i<100;i++ {
        go func() {
           p,v := ring.Read()
           fmt.Println(p,v)
        }()
    }


}