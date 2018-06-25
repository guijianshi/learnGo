package main

import "fmt"

func main()  {
	var a = 100
	testIf(a)
	testSwitch(a)
	testSelect()
}

func testIf(a int)  {
	if a < 10 {
		fmt.Println("小于20")
	} else if a < 20 {
		fmt.Println("{10, 20}")
	} else {
		fmt.Println("{20,无穷}")
	}
}

func testSwitch(a int)  {
	switch a {
	case 10:
		fmt.Println("值为10")
	case 20:
		fmt.Println("值为20")
	case 100:
		fmt.Println("值为100")
	default:
		fmt.Println("该值不支持")
	}
	fmt.Println("你的值是1")
}

func testSelect()  {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

