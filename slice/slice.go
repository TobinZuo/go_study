package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	// ======================================//
	// null，只定义，未初始化
	var s1 []int
	b1, _ := json.Marshal(s1)
	fmt.Println(string(b1))

	// []，初始化了
	s2 := []int{}
	b2, _ := json.Marshal(s2)
	fmt.Println(string(b2))

	// ======================================//
	// 减少边界检查，广泛运用于encoding/binary库，如果开头不做，编译成汇编指令就会每次都做编译检查
	//func (littleEndian) PutUint64(b []byte, v uint64) {
	//	_ = b[7] // early bounds check to guarantee safety of writes below
	//	b[0] = byte(v)
	//	b[1] = byte(v >> 8)
	//	b[2] = byte(v >> 16)
	//	b[3] = byte(v >> 24)
	//	b[4] = byte(v >> 32)
	//	b[5] = byte(v >> 40)
	//	b[6] = byte(v >> 48)
	//	b[7] = byte(v >> 56)
	//}

}