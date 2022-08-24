package main

import (
	"fmt"
	"math"
)

func main() {

	// int{num}

	var minInt8 int8 = math.MinInt8 // 1 byte == 8 bit
	var maxInt8 int8 = math.MaxInt8 // 1 byte == 8 bit
	fmt.Printf("int8: \n min: %d \n max: %d\n\n", minInt8, maxInt8)

	var minInt16 int16 = math.MinInt16 // 2 byte == 16 bit
	var maxInt16 int16 = math.MaxInt16 // 2 byte == 16 bit
	fmt.Printf("int16: \n min: %d \n max: %d\n\n", minInt16, maxInt16)

	var minInt32 int32 = math.MinInt32 // 4 byte == 32 bit
	var maxInt32 int32 = math.MaxInt32 // 4 byte == 32 bit
	fmt.Printf("int32: \n min: %d \n max: %d\n\n", minInt32, maxInt32)

	var minInt64 int64 = math.MinInt64 // 8 byte == 64 bit (x2 int32)
	var maxInt64 int64 = math.MaxInt64 // 8 byte == 64 bit (x2 int32)
	fmt.Printf("int64: \n min: %d \n max: %d\n\n", minInt64, maxInt64)

	//==============================================================

	// uint{num}

	var minUint8 uint8 = 0             // 1 byte == 8 bit
	var maxUint8 uint8 = math.MaxUint8 // 1 byte == 8 bit
	fmt.Printf("uint8: \n min: %d \n max: %d\n\n", minUint8, maxUint8)

	var minUint16 uint16 = 0              // 2 byte == 16 bit
	var maxUint16 uint16 = math.MaxUint16 // 2 byte == 16 bit
	fmt.Printf("uint16: \n min: %d \n max: %d\n\n", minUint16, maxUint16)

	var minUint32 uint32 = 0              // 4 byte == 32 bit
	var maxUint32 uint32 = math.MaxUint32 // 4 byte == 32 bit
	fmt.Printf("uint32: \n min: %d \n max: %d\n\n", minUint32, maxUint32)

	var minUint64 uint64 = 0              // 8 byte == 64 bit (x2 int32)
	var maxUint64 uint64 = math.MaxUint64 // 8 byte == 64 bit (x2 int32)
	fmt.Printf("uint64: \n min: %d \n max: %d\n\n", minUint64, maxUint64)

	//==============================================================

	// float{num}

	var exFloat32 float32 = 3.14 // 4 byte == 32 bit
	fmt.Printf("float32: %f\n\n", exFloat32)

	var exFloat64 float64 = 3.14 // 4 byte == 32 bit
	fmt.Printf("float64: %f\n\n", exFloat64)

	//==============================================================

	// other

	var minByte byte = 0             // synonym uint8
	var maxByte byte = math.MaxUint8 // synonym uint8
	fmt.Printf("byte: \n min: %d \n max: %d\n\n", minByte, maxByte)

	var minRune rune = math.MinInt32 // synonym int32
	var maxRune rune = math.MaxInt32 // synonym int32
	fmt.Printf("rune: \n min: %d \n max: %d\n\n", minRune, maxRune)

	var minInt int = math.MinInt // 4 byte or 8 byte
	var maxInt int = math.MaxInt // 4 byte or 8 byte
	fmt.Printf("int: \n min: %d \n max: %d\n\n", minInt, maxInt)

	var minUint uint = 0            // 4 byte or 8 byte
	var maxUint uint = math.MaxUint // 4 byte or 8 byte
	fmt.Printf("uint: \n min: %d \n max: %d\n\n", minUint, maxUint)

	var exString string = "Hello"
	fmt.Printf("string: %s\n\n", exString)

	var exBool bool = true
	fmt.Printf("bool: %t\n\n", exBool)

}
