package main

import (
	"fmt"
)

func main() {
	arg := []byte{1, 2, 3, 4, 5, 6, 7}
	arg2 := []byte{8, 9, 1}

	buf := make([]byte, 32)
	buf[0] = 0x01
	buf[1] = 0x02
	copy(buf[2:9], arg)
	copy(buf[9:12], arg2)

	for i := 12; i < 32; i++ {
		buf[i] = 0x20
	}

	fmt.Println(buf)
}
