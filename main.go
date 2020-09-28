package main

import "fmt"

func main() {
	var height uint16
	var width uint16
	fmt.Scan(&height)
	fmt.Scan(&width)
	area := height * width / 2.0
	fmt.Print(area)
}
