package control_flow

import (
	"fmt"
	"math/rand"
	"time"
)

func SimpleIf() {
	fmt.Print("Enter number 10: ")
	var num int
	fmt.Scan(&num)

	if num == 10 {
		fmt.Println("You entered correct number!")
	} else {
		fmt.Println("Do you know how number 10 is written?")
	}
}

func IfInitStatement() {
	fmt.Print("Enter number between 0 and 5: ")

	var guess int
	fmt.Scan(&guess)

	if secret := rand.Intn(5); secret == guess {
		fmt.Println("You guess is right!")
	} else {
		fmt.Println("Not today")
	}
}

func IfElseIfElse() {
	if h := time.Now().Hour(); h < 12 {
		fmt.Println("Now is AM time.")
	} else if h > 19 {
		fmt.Println("Now is evening time.")
	} else {
		fmt.Println("Now is afternoon time.")
	}
}
