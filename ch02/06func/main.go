package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op byte) (int, error) {
	switch op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		q, _, err := div(a, b)
		return q, err
	default:
		return 0, fmt.Errorf("不支持的运算：%c", op)
	}
}

// 带余数的除法
func div(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("除零")
	}
	return a / b, a % b, nil
}
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("调用函数 %s(%d, %d)\n", opName, a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func sum(nums ...int) (s int) {
	for _, i := range nums {
		s += i
	}
	return s
}
func main() {
	ops := []byte{'+', '-', '*', '/', '%'}
	a, b := 123, 78
	for _, op := range ops {
		result, err := eval(a, b, op)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%d %c %d = %d\n", a, op, b, result)
	}
	fmt.Println(div(123, 4))
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(func(a, b int) int {
		return a*a + b*b
	}, 3, 4))
	fmt.Println(sum(11, 22, 33, 44, 55, 66))
}
