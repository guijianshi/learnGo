package main

import "fmt"

func main()  {
	var a = 10

	fmt.Printf("变量的地址: %x\n", &a  )
	var b = 20
	swap(&a, &b)

	fmt.Printf("交换后a的值:%d", a)
	fmt.Printf("交换后b的值:%d", b)
}

func swap(x *int, y *int)  {
	*x, *y = *y, *x
}