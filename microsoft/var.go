package main

import "fmt"

// declaring and initializing variables with types. go will throw an error if you don't use variables.
var (
	ip4_address  string
	ipv6_address string
	ttl          int
)

// constants are static values
const (
	StatusOK              = 0
	StatusConnectionReset = 1
	StatusOtherError      = 2
)

func main() {
	ipv4_address, ipv6_address := "10.0.0.1", "FE80::"
	ttl := 255
	fmt.Println(ipv4_address, ipv6_address, ttl)
}
