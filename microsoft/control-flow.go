package main

import "fmt"

func main() {
	ttl := 255
	if ttl == 0 {
		fmt.Println(ttl, "value is 0 the packet is expired")
	}

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum of 1..100 is", sum)

}
