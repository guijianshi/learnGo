package main

import "fmt"

func main()  {
	testFor()
}

func testFor() int  {
	for a := 100; a > 0; a-- {
		fmt.Println(a)
	}
	b := true
	for b {
		fmt.Println("首次进入")
		b = false
	}

	numbers := [6]int{1, 2, 3, 5}
	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
	return 1
}



