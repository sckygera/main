package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Введите 2 целых числа в диапазане от -32768 до 32767 ")

	var number, number2 int64

	fmt.Scan(&number, &number2)

	result := int64(number) * int64(number2)
	fmt.Println(result)
	if result < 0 && result >= math.MinInt8 {
		fmt.Println("Результат можно сохранить в int8")
	} else if result < 0 && result < math.MinInt8 && result > math.MinInt16 {
		fmt.Println("Результат можно сохранить в int16")
	} else if result < 0 && result < math.MinInt16 && result > math.MinInt32 {
		fmt.Println("Результат можно сохранить в int32")
	}

	if result >= 0 && result <= math.MaxUint8 {
		fmt.Println("Результат можно сохранить в uint8")
	} else if result >= 0 && result <= math.MaxUint16 && result > math.MaxUint8 {
		fmt.Println("Результат можно сохранить в uint16")
	} else if result >= 0 && result <= math.MaxUint32 && result > math.MaxUint16 {
		fmt.Println("Результат можно сохранить в uint32")
	}

}
