package main

import (
	"math/rand"
	"time"
)

// 对arr[l...r]部分进行partition操作
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func partition (arr []int,l,r int) int {
	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	rand.Seed(time.Now().Unix())
	randIndex := rand.Int()%(r-l+1)+l
	arr[l],arr[randIndex] = arr[randIndex],arr[l]

	v := arr[l]
	j := l
	// 先把小于随机点的数移动到左边，记录移动的次数，最后根据移动的次数把随机点交换到移动次数的位置，保证右边是大于这个随机数的
	for i := l+1; i<=r; i++ {
		if arr[i] < v {
			j++
			arr[j],arr[i] = arr[i],arr[j]
		}
	}
	arr[l],arr[j] = arr[j],arr[l]
	return j
}

func QuickSort(arr []int,l,r int)  {
	if l >= r {
		return
	}
	p := partition(arr,l,r)
	QuickSort(arr,l,p-1)
	QuickSort(arr,p+1,r)
}

