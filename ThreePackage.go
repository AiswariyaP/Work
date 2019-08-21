package main

import (
	"fmt"

	"D:/Training/goLang_practical/Golang_practical/src/go/language/exporting/try/counters"
)

func main() {

	// Create a variable of the exported type and initialize the value to 10.
	counter1 := counters.AlertCounter1(50)

	fmt.Printf("Counter: %d\n", counter1)

	counter2 := counters.AlertCounter2(14)

	fmt.Printf("Counter: %d\n", counter2)

	counter3 := counters.AlertCounter3(50)

	fmt.Printf("Counter: %d\n", counter3)
}