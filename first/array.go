package main

import "fmt"

var fruit [3] string //定义数组
var balance = [5]float32{100.0, 3.0, 4.0, 5.0, 3.3} // 定义并初始化数组
var multi_arr = [3][4]int{
	{0, 1, 2, 3},
	{1, 2, 3, 4},
	{2, 3, 4, 5},
}

func main() {
	fruit[0] = "apple"
	fmt.Println(balance[1])
	fmt.Println(fruit[0])
	fmt.Println(multi_arr[1][2])
	arr_1 := []int{1,4,6,7,8}
	fmt.Println(getAverage(arr_1, 5))
}

func getAverage(arr []int, size int) float32 {
	var i,sum int
	var avg float32

	for i = 0; i < size;i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg;
}