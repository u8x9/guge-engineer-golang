package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 999
}
func basic() {
	arr := [...]int{11, 22, 33, 44, 55, 66, 77, 88}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:](s1) = ", s1)
	s2 := arr[:]
	fmt.Println("arr[:](s2) = ", s2)

	fmt.Println("--- after update s1 ---")
	updateSlice(s1)
	fmt.Println("s1 = ", s1)
	fmt.Println("arr = ", arr)

	fmt.Println("--- after update s2 ---")
	updateSlice(s2)
	fmt.Println("s2 = ", s2)
	fmt.Println("arr = ", arr)

	fmt.Println("--- reslice ---")
	s2 = s2[:5]
	fmt.Println("s2 = ", s2)
	s2 = s2[2:]
	fmt.Println("s2 = ", s2)
}

// slice 可以向后扩展，但不能向前扩展
// 取下标由 len 决定
// reslice  由 cap 决定
func extend() {
	arr := [...]int{0, 11, 22, 33, 44, 55, 66, 77}
	s1 := arr[2:6] // [22, 33, 44, 55]
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n",
		s1, len(s1), cap(s1))
	s2 := s1[3:5] // [55, 66] ?
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n",
		s2, len(s2), cap(s2))
}
func main() {
	// basic()
	extend()
}
