package ringbuffer

import (
    "testing"
    "time"
    "fmt"
)

func TestDateSecond(t *testing.T) {
    s := getFiveSecond()
    fmt.Println(s)
    time.Sleep(1*time.Second)

    s = getFiveSecond()
    fmt.Println(s)
    time.Sleep(1*time.Second)

    s = getFiveSecond()
    fmt.Println(s)
    time.Sleep(1*time.Second)

    s = getFiveSecond()
    fmt.Println(s)
    time.Sleep(1*time.Second)

    s = getFiveSecond()
    fmt.Println(s)
    time.Sleep(1*time.Second)

    s = getFiveSecond()
    fmt.Println(s)
    time.Sleep(1*time.Second)

    s = getFiveSecond()
    fmt.Println(s)
    time.Sleep(1*time.Second)

    s = getFiveSecond()
    fmt.Println(s)
    time.Sleep(1*time.Second)
}

func getFiveSecond()  int64 {
    now := time.Now()

    second := now.Second()
    remainder := second % 5
    if remainder == 0 {
        return now.Unix()
    }

    return int64(5 - remainder) + now.Unix()
}
