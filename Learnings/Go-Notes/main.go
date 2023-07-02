package main

import (
	"fmt"
	"time"
)

/*
*		refer : https://gobyexample.com/
 */

func main() {
	fmt.Println("Hello")

	/*
		Strings can be added using + operator

		In Go variables are declared explicitly, and used by compiler to eg: to check type correctness
		of function calls
	*/

	var a = "initial"
	fmt.Println(a)
	var b, c = 1, 2 // can be declared in a single line
	fmt.Println(b, c)

	/*
		Go infers the type of initialized variable.
		Variables declared without a corresponding initialization are zero valued.
		eg : the zero value for int is 0

		:= -> is shorthand for declaring and initializing a variable.


		Constants:
			const s string = "constant"
		Once a variable is declared const it cannot be changed later

	*/

	const s string = "constant"
	fmt.Println(s)

	/*
		Loops
	*/

	//for loop
	for i := 0; i < len("string"); i++ {
		fmt.Println(i)
	}

	for i := range "string" {
		fmt.Println(i)
	}

	for k, v := range "string" {
		fmt.Println(k, v)
		break // to break in between
	}

	//equivalent to while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// conditionals
	if true {

	} else if false {

	} else {
		// what the hell is this condition :)
	}

	// switch statements
	switch i {
	case 1:
		fmt.Println(i)
	case 2:
		fmt.Println(i)
	}

	t := time.Now()
	switch {
	case t.Hour() < 2:
		fmt.Println("midnight")
	default:
		fmt.Println("Get up")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("week day")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("not sure %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI(nil)
}
