package main

import (
	"fmt"
)

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int)  {
	n *= 2	
}

func doublePtr(n *int)  {
	*n *= 2	
}

func dd(n int) int {
	return n * 3
}

func main()  {
	values :=[]int{1, 2, 3, 4}
	doubleAt(values, 3)
	doubleAt(values, 0)
	fmt.Println(values)

	val := 10
	double(val)
	fmt.Println(val)
	b := dd(val)
	fmt.Println(b)
	doublePtr(&val)
	fmt.Println(val)
	dd(val)
	fmt.Println(dd(val))
}