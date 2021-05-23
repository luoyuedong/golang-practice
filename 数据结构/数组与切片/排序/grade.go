package main

func GradeSort(grade []int) {
	// 初始化出试卷范围[0-100]
	arr := make([]int,101)
	for _,v := range grade {
		// 每一个成绩装到对应的桶中
		arr[v]++
	}
	// 把成绩装回去
	index := 0
	for i,v := range arr {
		// 桶空了就继续下一个
		if v == 0 {
			continue
		}else {
			// 如果桶不为空,也就是说arr[i]不为空
			// 意思就是成绩为 i 的人有多少个
			// 所以我们把这些成绩装回grade数组中去
			for v != 0 {
				v--
				grade[index] = i
				index++
			}
		}
	}
}

