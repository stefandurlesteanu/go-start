package main

import "fmt"

func main() {
	// ----- IF ----
	statePopulation := map[string]int{
		"Texas": 1234,
		"LA":    2345,
		"NY":    3456,
	}

	// Simple IF
	if 1 > 0 {
		fmt.Println("Always true!")
	}

	// Initializer IF statement
	if pop, ok := statePopulation["Texas"]; ok {
		fmt.Println(pop)
	}

	// IF ELSE IF ELSE
	if statePopulation["Texas"] == 112 {
		fmt.Println("That is not true")
	} else if statePopulation["NY"] == 3456 {
		fmt.Println("That is true")
	} else {
		fmt.Println("Never get here")
	}

	// ----- SWITCH -----
	switch i := 2 + 3; i {
	case 1, 5:
		fmt.Println("One or Five")
	case 2, 4, 6:
		fmt.Println("Two, Four or Six")
	default:
		fmt.Println("Not in range")
	}

	//fallthrough && break
	j := 10
	switch {
	case j <= 10:
		fmt.Println("Less of equal to ten")
		fallthrough
	case j <= 20:
		fmt.Println("Less or equal to twenty")
		if j == 20 {
			break
		}
		fmt.Println("This won't print")
	default:
		fmt.Println("Default")
	}
}
