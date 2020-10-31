package main

import (
	"fmt"
)

func main()  {
	nums := []int{64, 8, 42, 4, 23, 5, 105}
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	fmt.Println(max)
}