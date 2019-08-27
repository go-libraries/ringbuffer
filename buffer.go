package ringbuffer

import (
    "sync"
    "errors"
)

type RingBuffer struct {
    rw *sync.RWMutex
    buffer []interface{} //cache buffer
    size uint64     //buffer size
    r uint64        //read  pos
    w uint64        //write pos
    d uint64        //done  pos
}

func MakeRingBuffer(size uint64) *RingBuffer {
    return &RingBuffer{
        rw:new(sync.RWMutex),
        size:size,
        buffer:make([]interface{}, size),
    }
}

//check buffer full
func (ring *RingBuffer) IsFull() bool {
    if ring.w == ring.r {
        return false
    }

    if (ring.w + 1) % ring.size == ring.r {
        return true
    }

    return false
}

func (ring *RingBuffer) IsEmpty() bool {
    if ring.w == ring.r {
        return true
    }

    return false
}

//loop write in buffer
func (ring *RingBuffer) Write(i interface{})  error {
    ring.rw.Lock()
    defer ring.rw.Unlock()

    if ring.IsFull() {
        return errors.New("it's full")
    }

    ring.buffer[ring.w] = i
    if (ring.w + 1) == ring.size  {
        ring.w = 0
    } else {
        ring.w += 1
    }

    return nil
}

//read the last value
func (ring *RingBuffer) Read()  (uint64,interface{}){
    ring.rw.RLock()
    defer ring.rw.RUnlock()

    if ring.IsEmpty() {
        return ring.r, nil
    }


    now := ring.r
    next := ring.r+1
    if next == ring.size  {
        next = 0
    }
    v := ring.buffer[now]
    ring.r = next

    return now, v
}

// rewrite the pos value
//
// pos,err := ring.Read()
// if err != nil { xxx }
// defer ring.Done(pos)
//
func (ring *RingBuffer) Done(pos uint64) error{
    v := ring.buffer[pos]

    if v == nil {
        return errors.New("not exists value")
    }

    ring.buffer[pos] = nil
    return nil
}



