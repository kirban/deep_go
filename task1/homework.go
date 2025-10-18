package main

import (
	"fmt"
	"math"
)

func ToLittleEndian(number uint32) uint32 {
	const base = 1 << 8

	var result uint32

	mod := number % base
	div := number / base

	for i := 4; i >= 0; i-- {
		result += mod * uint32(math.Pow(float64(base), float64(i-1)))
		mod = div % base
		div = div / base
	}

	return result
}

func main() {
	var input uint32 = 0x010203FF

	fmt.Printf("%#08X\n", ToLittleEndian(input))
}
