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
func appendElem() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10) // 10 替换了最后的 7
	fmt.Printf("s3=%v, arr=%v\n", s3, arr)
	s4 := append(s3, 11) // 新的底层数组
	fmt.Printf("s4=%v, arr=%v\n", s4, arr)
	s5 := append(s4, 12)
	fmt.Printf("s5=%v, arr=%v\n", s5, arr)
}
func sliceInfo(s []int) {
	fmt.Println("len(s) =", len(s), ", cap(s) =", cap(s), " ==>", s)
}
func sliceOps() {
	fmt.Println(" -- create slice --- ")
	var s []int
	for i := 0; i < 100; i++ {
		sliceInfo(s)
		s = append(s, 2*i+1)
	}

	s1 := []int{2, 4, 6, 8}
	sliceInfo(s1)

	s2 := make([]int, 16)
	sliceInfo(s2)

	s3 := make([]int, 10, 32)
	sliceInfo(s3)

	fmt.Println(" -- copy slice --- ")
	copy(s2, s1)
	sliceInfo(s2)

	fmt.Println(" -- delete slice element --")
	// 删除 s2[3]
	s2 = append(s2[:3], s2[4:]...)
	sliceInfo(s2)

	fmt.Println(" -- pop --")
	front := s2[0]
	s2 = s2[1:]
	sliceInfo(s2)
	fmt.Println("pop: ", front)

	fmt.Println(" -- unshift --")
	lastIdx := len(s2) - 1
	tail := s2[lastIdx]
	s2 = s2[:lastIdx]
	sliceInfo(s2)
	fmt.Println("unshift: ", tail)

}
func main() {
	// basic()
	// extend()
	// appendElem()
	sliceOps()
}
