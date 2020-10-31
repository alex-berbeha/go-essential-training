package main

import (
	"fmt"
)

func main()  {
	book := "The colour of magic"
	fmt.Println(book)
	
	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// strings in go are immutable
 
	fmt.Println(book[4:11])

	fmt.Println(book[:4])

	fmt.Println(book[4:])

	fmt.Println("t" + book[1:])

	fmt.Println("It was 0.5 price!")

	poem := `
The road goes ever on
Down from the door where it began
...
	`

	fmt.Println(poem)

}