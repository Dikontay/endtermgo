package internal

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func Start() {
	fmt.Println("HELLO PLEASE CHOOSE THE TYPE OF ANIMAL")
	fmt.Println("1. DOG 2. Owl 3. BEAR ")

	var option string

	_, err := fmt.Scan(&option) //Getting number of animal

	handleErrors(err)
	err = check(option)
	if err != nil {
		fmt.Println(err) //Checking if number of that animal exists
	}

	animalBehavior := createAnimal(option)
	//Creating Animal Object
	if animalBehavior == nil {
		fmt.Println("Failed to create animal")
		return
	}

	// Create an instance of Animal and set its behavior
	animalContext := &AnimalContext{ //Wrapping animal object into context
		Behavior: animalBehavior,
	}
	timer := &Timer{
		secondSinceStart:   time.Now(),
		animal: animalBehavior,
	}

	go timer.updateState()
	Keycontrol(animalContext)


}

func createAnimal(num string) AnimalStaregy {

	var behavior AnimalStaregy

	switch num {
	case "1":
		behavior = &Dog{Clean: 10, Food: 10, Mood: 10, TimeToBeDirty: 4, TimeToEat: 6, TimeToPlay: 3}
	case "2":
		behavior = &Owl{Clean: 10, Food: 10, Mood: 10, TimeToBeDirty: 10, TimeToEat: 12, TimeToPlay: 12}
	case "3":
		behavior = &Bear{Clean: 10, Food: 10, Mood: 10, TimeToBeDirty: 20, TimeToEat: 2, TimeToPlay: 6}
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
