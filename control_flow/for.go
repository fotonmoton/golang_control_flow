package control_flow

import (
	"fmt"
	"math/rand"
)

func ForSimple() {

	// All the same
	// for ; true; {
	// }
	// for true {
	// }
	// for ; ; {
	// }

	for {
		fmt.Println("Let's burn thees CPU cycles!!!!111")
	}
}

func ForWithoutCondition() {
	for j := 0; ; j++ {
		if j == 10 {
			fmt.Printf("We found it! %d\n", j)
			break
		}
		fmt.Println(j)
	}
}

func ForIterator() {
	for iterator := 1; iterator < 100; iterator += iterator * 2 {
		fmt.Printf("We can go up %d\n", iterator)
	}

	for iterator := 10; iterator > 0; iterator-- {
		fmt.Printf("And down %d\n", iterator)
	}
}

func ForConditionOnly() {
	bigStep := 0
	for bigStep < 100 {
		fmt.Printf("As simple as that! %d\n", bigStep)
		bigStep += 10
	}
}

func ForExpanded() {
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func ForContinue() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func ForContainer() {

	countries := map[string]string{"Ukraine": "Kyiv", "Finland": "Helsinki", "Sweden": "Stockholm", "Australia": "Canberra"}

	for country, capital := range countries {
		fmt.Printf("%s is the capital of %s\n", capital, country)
	}

	numbers := []int{}
	for i := 0; i < 10; i++ {
		numbers = append(numbers, rand.Intn(10))
	}

	for index, number := range numbers {
		fmt.Printf("index %d has a %d number\n", index, number)
	}

	dumbCount := 0
	for range numbers {
		dumbCount++
	}

	if len(numbers) == dumbCount {
		fmt.Printf("Yes, len(numbers) == dumbCount == %d. But you can use len function, you dumb!", dumbCount)
	} else {
		fmt.Println("Hm...")
	}

	var nilMap map[int]string = nil

	for range nilMap {
		fmt.Println("Unreachable code!")
	}

	for country := range countries {
		if country != "Ukraine" {
			fmt.Printf("Goodbye, %s...\n", country)
			delete(countries, country)
		}
	}

	fmt.Println("Countries that are left:")
	for country, capital := range countries {
		fmt.Printf("%s is the capital of %s\n", capital, country)
	}
}

func ForCopies() {
	type Car struct {
		brand   string
		founded int
	}

	oldCars := [2]Car{{brand: "Buick", founded: 1899}, {brand: "Cadillac", founded: 1902}}
	for i, car := range oldCars {
		fmt.Println(i, car)

		// No effect for the ranged array
		oldCars[1].brand = "BMW"

		// This is a copy, so no effect for the ranged array and no effect for the loop body
		car.founded = 2023
	}
	fmt.Println("oldCars:", &oldCars)

	oldCarsSlice := []Car{{brand: "Buick", founded: 1899}, {brand: "Cadillac", founded: 1902}}
	for i, car := range oldCarsSlice {
		fmt.Println(i, car)

		// This modification has effect for the loop body and for the original slice
		oldCarsSlice[1].brand = "BMW"

		// This stays as before because of a copy
		car.founded = 2023
	}
	fmt.Println("oldCarsSlice:", &oldCarsSlice)
}
