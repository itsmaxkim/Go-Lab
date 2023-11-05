package main

import "fmt"

func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Vancouver", "Kamloops", "Kelowna":
		region, continent = "British Columbia", "Canada"
	case "Toronto", "Vaughan", "Markham":
		region, continent = "Ontario", "Canada"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}
func main() {
	region, continent := location("Vancouver")
	fmt.Printf("Max works in %s, %s\n", region, continent)
}
