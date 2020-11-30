package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":    "u8x9",
		"course":  "golang",
		"site":    "github.com",
		"quality": "sogood",
	}
	fmt.Println(m)

	m2 := make(map[string]int) // empty map
	fmt.Println(m2)

	var m3 map[string]bool // nil
	fmt.Println(m3)

	fmt.Println("-- 遍历 --")
	for k, v := range m {
		fmt.Printf("m[%s] = %s\n", k, v)
	}

	fmt.Println("-- 获取元素 --")
	courseName := m["course"]
	fmt.Println("course:", courseName)
	// notExistsElement := m["not_exists"]
	// fmt.Printf("not exists element: %#v\n", notExistsElement)
	if notExistsElement, ok := m["not_exists"]; ok {
		fmt.Printf("not exists element: %#v\n", notExistsElement)
	} else {
		fmt.Println("but the element no exists.")
	}

	fmt.Println("-- 删除元素 --")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
