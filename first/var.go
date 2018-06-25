package main

import "fmt"

var name string  = "lin" // 声明变量类型并赋值
var sex  = "男" // 声明变量并赋值(根据值自动分配类型)
var age, phone int // 多变量声明
var c , d = 1, 2
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

const TYPE_1 int = 1
const TYPE_2  = 2
const TYPE_3, TYPE_4  = 3, 4

const (
	UNKNOWN = iota // 在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1
	FEMALE
	MALE
)

func main() {
	city  := "北京" // 声明并赋值
	age = 18
	phone = 1526372632
	a, b = 1, false
	e := &b
	*e = true
	fmt.Println(a, b, c, d)
	fmt.Println(TYPE_1, TYPE_2, TYPE_3, TYPE_4, UNKNOWN, FEMALE, MALE)
	fmt.Printf("姓名:%s,性别:%s,年龄%d,电话%d,籍贯:%s", name, sex, age, phone,city)
}