package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61,
	}

	fmt.Println(len(stocks))
	fmt.Println("---------------")
	fmt.Println(stocks["MSFT"])
	fmt.Println("---------------")
	fmt.Println(stocks["TSLA"])
	fmt.Println("---------------")
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
		fmt.Println(ok)
	} else { 
		fmt.Println(value)
	}
	fmt.Println("---------------")
	stocks["TSLA"] = 322.12

	fmt.Println("---------------")
	delete(stocks, "AMZN")
	fmt.Println(stocks)
	
	fmt.Println("---------------")
	for key := range stocks {
		fmt.Println(key)
	}


	fmt.Println("---------------")
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}

	fmt.Println("---------------")
	for _, value := range stocks {
		fmt.Println(value)
	}
}
