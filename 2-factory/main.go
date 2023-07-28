package main

import (
	"fmt"
)

func main() {
	tiger, _ := getAnimal("bengala")
	leon, _ := getAnimal("africano")

	printDetails(tiger)
	printDetails(leon)
}

func printDetails(a IAnimal) {
	fmt.Printf("Name: %s", a.getName())
	fmt.Println()
	fmt.Printf("Race: %s", a.getRace())
	fmt.Println()
	fmt.Printf("Danger: %d", a.getDanger())
	fmt.Println()
}
