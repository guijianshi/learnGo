package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(max(10, 20))
	surname, name := getName()
	fmt.Println(surname)
	fmt.Println(name)
	status1 := 2
	status2 := 2
	changeStatus(status1, &status2)
	fmt.Printf("status1:%d,status2:%d\n", status1, status2)
	getSquareRoot()
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func getName() (string, string)  {
	return "lin", "yang"
}

func changeStatus(status1 int, status2 *int) {
	status1--
	*status2--
}

func getSquareRoot()  {
	getSquareRootFunc := func(x float64) float64 {
		return  math.Sqrt(x)
	}
	fmt.Println(getSquareRootFunc(9))
}

func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}

