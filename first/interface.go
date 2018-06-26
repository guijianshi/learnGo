package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {

}

func (nokiaPhone NokiaPhone) call()  {
	fmt.Println("我正在用诺基亚给你打电话")
}

type IPhone struct {

}

func (iPhone IPhone) call() {
	fmt.Println("我正在用苹果给你打电话")
}

func main() {
	var phone Phone
	phone = new (NokiaPhone)
	phone.call()
	phone = new (IPhone)
	phone.call()
}
