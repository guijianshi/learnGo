package main

import "fmt"

func main() {
	fmt.Println(fibonacci(30))
	fmt.Println(factorial(30))
}

/* 斐波那契数列 */
func fibonacci(a int) (result int)  {
	if a > 2 {
		return fibonacci(a - 1) + fibonacci(a - 2)
	} else {
		return  a;
	}
}

/* 阶乘 */
func factorial(n uint64) (result uint64)  {
	if n > 0 {
		return  n * factorial(n - 1)
	} else  {
		return 1;
	}
}