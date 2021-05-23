package main

import (
	"fmt"
	"strconv"
	"strings"
)

//  转换int整形为数组
func Digits(n int) []int {
	s := strconv.Itoa(n)
	d := make([]int, len(s))
	for i, l := range strings.Split(s, "") {
		d[i], _ = strconv.Atoi(l)
	}
	return d
}

func DigitsE(n int) []int {
	s := strconv.Itoa(n)
	d := make([]int, len(s))
	for k, v := range strings.Split(s, "") {
		d[k], _ = strconv.Atoi(v)
	}
	return d
}
func main() {
	r := Digits(123) // [1 2 3]
	fmt.Println(r)
}
