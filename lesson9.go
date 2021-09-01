package main

import (
	"fmt"
)

func main() {

	const MaxUint16 uint16 = 1<<16 - 1
	const MaxUint32 uint32 = 1<<32 - 1
	const MaxUint8 uint8 = 1<<8 - 1

	for i := MaxUint32; i != 0; i++ {
		overflow8 := i / (uint32(MaxUint8) + uint32(1))
		overflow16 := i / (uint32(MaxUint16) + uint32(1))

		fmt.Printf("переполнений типов uint8: %d. Переполнений типа uint16: %d. ", overflow8, overflow16)
	}
}
