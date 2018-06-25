package main

import "fmt"

type Book struct {
	title string
	author string
	subject string
	price float32
	book_id int
}

func main()  {
	var Book1 Book
	var Book2 Book
	Book1.title = "论持久战"
	Book1.author = "毛泽东"
	Book1.subject = "毛泽东写的论持久战很不错"
	Book1.price = 18.82
	Book1.book_id = 1

	Book2.title = "资本论"
	Book2.author = "马克思"
	Book2.subject = "马克思成名之作"
	Book2.price = 28.8
	Book2.book_id = 2

	/* 打印 Book1 信息 */
	fmt.Printf( "Book 1 title : %s\n", Book1.title)
	fmt.Printf( "Book 1 author : %s\n", Book1.author)
	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
	fmt.Printf( "Book 1 price : %0.2f\n", Book1.price)
	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf( "Book 2 title : %s\n", Book2.title)
	fmt.Printf( "Book 2 author : %s\n", Book2.author)
	fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
	fmt.Printf( "Book 2 price : %0.2f\n", Book2.price)

	fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
	printBookInfo(&Book1)
}

func printBookInfo(book *Book)  {
	/* 打印 Book1 信息 */
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book price : %0.2f\n", book.price)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
}