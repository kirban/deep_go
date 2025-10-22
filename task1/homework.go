package main

import (
	"fmt"
	"unsafe"
)

func ToLittleEndian(number uint32) uint32 {
	inputPtr := unsafe.Pointer(&number)
	inputSize := int(unsafe.Sizeof(number))

	var result uint32
	outPtr := unsafe.Pointer(&result)

	for i := 0; i < inputSize; i++ {
		*(*int8)(unsafe.Add(outPtr, inputSize-i-1)) = *(*int8)(unsafe.Add(inputPtr, i))
	}

	return result
}

func main() {
	var input uint32 = 0x010203FF

	fmt.Printf("%#08X\n", ToLittleEndian(input))
}
