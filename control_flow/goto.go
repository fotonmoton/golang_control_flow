package control_flow

import "fmt"

func GoToSimple() {
	i := 0

Next: // here, a label is declared.
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next // execution jumps
	}
}

func GoToAfter() {

	fmt.Println("Your learning about 'goto' statement")
	// We create a for loop which runs until i is 10
	for i := 0; i < 10; i++ {
		fmt.Printf("Index: %d\n", i)
		if i == 5 {
			// When i is 5, lets exit by using goto
			goto exit
		}
	}
	fmt.Println("Skip this line here")
	// Create the exit label and insert code that should be executed when triggered
exit:
	fmt.Println("We are now exiting the program")
}

func GotToExitInMethodAbove() {
	// goto exit
}

func GotToLabels() {
	// goto Label1 // error
	// {
	// Label1:
	// 	goto Label2 // error
	// }
	// {
	// Label2:
	// }
}

func GoToJumpOverDeclaration() {
	// 	i := 0
	// Next:
	// 	if i >= 5 {
	// 		// error: jumps over declaration of k
	// 		goto Exit
	// 	}

	// 	k := i + i
	// 	fmt.Println(k)
	// 	i++
	// 	goto Next

	// 	// This label is declared in the scope of k,
	// 	// but its use is outside of the scope of k.

	// Exit:
}

func GoToJumpOverDeclarationFixed() {
	i := 0
Next:
	if i >= 5 {
		// no errors now
		goto Exit
	}

	// New scoped block
	{
		k := i + i
		fmt.Println(k)
	}
	i++
	goto Next

	// This label is declared in the scope of k,
	// but its use is outside of the scope of k.

Exit:
}

func GoToBreakLabel() {
	row := 0
Rows:
	for {
		col := 0
	Columns:
		for {

			if row == 10 {
				break Rows
			}

			if col == 10 {
				break Columns
			}
			fmt.Print(col)
			col++
		}
		fmt.Println()
		row++
	}
}

func GoToContinueLabel() {
	fmt.Println("Let's start cycling")
	i := 0
Outer:
	for {
		switch i {
		case 11:
			break Outer
		case 1, 5, 9:
			fmt.Println("Skipping this guys")
			i += 2
			continue Outer
		default:
			fmt.Println("Not yet")
		}
		i++
	}
	fmt.Println("Finish")
}
