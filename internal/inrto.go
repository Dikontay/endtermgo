package internal

import (
	"errors"
	"fmt"
	"os"
)

func Start() {
	fmt.Println("HELLO PLEASE CHOOSE THE TYPE OF ANIMAL")
	fmt.Println("1. DOG 2. Owl 3. BEAR ")

	var option string

	_, err := fmt.Scan(&option)

	handleErrors(err)
	err = check(option)
	if err != nil {
		fmt.Println(err)
	}

	animalBehavior := createAnimal(option)
	if animalBehavior == nil {
		fmt.Println("Failed to create animal")
		return
	}

	// Create an instance of Animal and set its behavior
	animalStaregy := &AnimalContext{
		behavior: animalBehavior,
	}
	animalStaregy.behavior.Print()
	// for {
		
	// }

	}



func createAnimal(num string) AnimalStaregy {

	var behavior AnimalStaregy

	switch num {
	case "1":
		behavior = &Dog{Name: "dog"}
	case "2":
		behavior = &Owl{Name: "owl"}
	case "3":
		behavior = &Bear{Name: "bear"}
	// Add cases for other animal types
	default:
		return nil
	}

	return behavior
}

func check(option string) error {
	if option == "1" || option == "2" || option == "3" || option == "4" {
		return nil
	}

	return errors.New("Incorrect option choosen")
}

func handleErrors(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
