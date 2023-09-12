package control_flow

import (
	"fmt"
	"math/rand"
)

func SwitchBasicCase() {

	// Compared with "=="
	switch n := rand.Intn(100); n % 9 {
	case 0:
		fmt.Println(n, "is a multiple of 9.")

		// Different from many other languages,
		// in Go, the execution will automatically
		// jumps out of the switch-case block at
		// the end of each branch block.
		// No "break" statement is needed here.
	case 1, 2, 3:

		fmt.Println(n, "mod 9 is 1, 2 or 3.")
		// Here, this "break" statement is nonsense.
	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
	// case 6, 7, 8:
	// The above case line might fail to compile,
	// for 6 is duplicate with the 6 in the last
	// case. The behavior is compiler dependent.
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}
}

func SwitchFallthrough() {
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		// This will not compile
		// fallthrough

		fmt.Println("n =", n)
		// The "fallthrough" statement makes the
		// execution slip into the next branch.
		fallthrough
	case 5, 6, 7, 8:
		// A new declared variable also called "n",
		// it is only visible in the current
		// branch code block.

		n := 99
		fmt.Println("n =", n) // 99
		fallthrough
	default:
		// This "n" is the switch expression "n".
		fmt.Println("n =", n)
		// This will not compile
		// fallthrough
	}
}

func SwitchBranchesAreOptional() {
	switch n := 5; n {
	}

	switch 5 {
	}

	switch _ = 5; {
	}

	switch {
	}
}

func SwitchWithoutStatementAndOperand() {
	switch { // <=> switch true {
	case true:
		fmt.Println("hello")
		// fallthrough
	default:
		fmt.Println("bye")
	}
}

func SwitchDefaultOrder() {
	switch n := rand.Intn(3); n {
	case 0:
		fmt.Println("n == 0")
	case 1:
		fmt.Println("n == 1")
	default:
		fmt.Println("n == 2")
	}

	switch n := rand.Intn(3); n {
	default:
		fmt.Println("n == 2")
	case 0:
		fmt.Println("n == 0")
	case 1:
		fmt.Println("n == 1")
	}

	switch n := rand.Intn(3); n {
	default:
		fmt.Println("n == 2")
	case 0:
		fmt.Println("n == 0")
	case 1:
		fmt.Println("n == 1")
	}
}
