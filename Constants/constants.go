package main

import "fmt"

const Pi = 3.14

func main() {
	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
	/**
	Hello 世界
	Happy 3.14 Day
	Go rules? true
	*/
}
