package main

import "fmt"

func main() {
	slice := make([]int,0)
	for i := 0 ;i < 10; i++ {
		// 动态的对切片进行扩容
		slice = append(slice, i)
	}
	fmt.Println(slice)

	// 调用PrintArr
	PrintArr(slice)
	// 看看是否切片的第一个元素改变了?
	fmt.Println(slice)
}

func PrintArr(arr []int) {
	arr[0] = 100
	for index, value := range arr {
		fmt.Printf("索引：%d, 值： %d\n",index,value)
	}
}
