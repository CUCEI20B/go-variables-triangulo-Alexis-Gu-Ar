package main

import "fmt"

func main() {
	var height uint64
	var width uint64
	fmt.Scan(&height)
	fmt.Scan(&width)
	area := height * width / 2
	fmt.Print(area)
}
