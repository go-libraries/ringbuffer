package ringbuffer

import (
    "testing"
    "fmt"
)

func TestSliceCopy(t *testing.T)  {
    slice := []int{1, 2, 3, 4, 5}
    sliceModify(&slice)
    fmt.Println(slice)
    sliceModifyValue(slice)
    fmt.Println(slice)

    sliceModifyValueExists(slice)
    fmt.Println(slice)
}

func sliceModify(slice *[]int) {
    *slice = append(*slice, 6)
}

func sliceModifyValue(slice []int) {
    slice = append(slice, 7)
}

func sliceModifyValueExists(slice []int) {
    slice[0] = 99
}
