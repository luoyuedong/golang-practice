package main

import "fmt"

func main() {
	// 桶排序
	grade := []int{4,66,66,67,55,55,66,99,100,4,67}
	GradeSort(grade)
	fmt.Println("桶排序的结果", grade)

	var arr []int
	for i:=9 ; i>=0;i--{
		arr = append(arr,i)
	}
	fmt.Println(arr)
	QuickSort(arr,0,len(arr)-1)
	fmt.Println("快速排序的结果", arr)
}

