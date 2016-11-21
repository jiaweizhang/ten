// _Switch statements_ express conditionals across many
// branches.

package main

import "fmt"
import "time"

func main() {

	// Here's a basic `switch`.
	i := 4
	fmt.Print("write ", i, " as ")
	switch i { // note that no breaks needed between cases and scopes are different
	case 1:
		fmt.Println("one")
		k := 15
		fmt.Println(k)
	case 2:
		fmt.Println("two")
		k := 15
		fmt.Println(k)
	case 3:
		fmt.Println("three")
	case 4:
	case 5:
		fmt.Println("4 or 5")
	}

	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}

// todo: type switches
