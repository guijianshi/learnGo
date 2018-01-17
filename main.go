package main

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	//var name string
	//var money float32
	//const LENGTH int = 10
	//age := 27
	//money = 0.232
	//name = "Hello World!"
	//vname1, vname2, vname3 := '1', '2', '3'
	//fmt.Printf("%s\n", name)
	//fmt.Printf("%f\n", money)
	//fmt.Printf("%d\n", age)
	//fmt.Printf("%s %s %s", vname1, vname2, vname3)
	//fmt.Printf("%d\n", LENGTH)
	const (
		a = "test"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	//fmt.Println(c)
	//
	//var Book1 Book
	//Book1.title = "深入面向对象"
	//Book1.author = "php"
	//Book1.subject = "深入面向对象PHP"
	//Book1.book_id = 1
	//
	//printBook(Book1)
	//
	//var p_bool *Book
	//p_bool = &Book1
	//fmt.Println(p_bool.author)
	//
	//var i int
	//for i = 0; i < 10; i++ {
	//	fmt.Printf("%d\t", fibonacci(i))
	//}
	//fmt.Println()
	test()
}

func printBook(book Book) {
	fmt.Printf("Book 1 title : %s\n", book.title)
	fmt.Printf("Book 1 author : %s\n", book.author)
	fmt.Printf("Book 1 subject : %s\n", book.subject)
	fmt.Printf("Book 1 book_id : %d\n", book.book_id)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func test() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("START Goroutines")

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c\n ", char)
			}
		}

	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\n Terminating Program")
}
