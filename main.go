package main

import (
	"fmt"
)

func main() {
	dog := [11][15]rune{
		{' ', ' ', '▉', '▉', '▉', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', '▉', ' ', ' ', ' ', '▉', '▉', ' ', ' ', ' ', ' ', ' ', ' ', '▉', '▉'},
		{' ', '▉', ' ', '▉', ' ', ' ', ' ', '▉', ' ', ' ', ' ', ' ', '▉', ' ', '▉'},
		{'▉', ' ', ' ', ' ', ' ', '▉', ' ', ' ', '▉', ' ', ' ', '▉', ' ', '▉', ' '},
		{'▉', '▉', ' ', ' ', ' ', '▉', ' ', ' ', '▉', '▉', '▉', '▉', '▉', ' ', ' '},
		{'▉', ' ', ' ', ' ', '▉', ' ', '▉', '▉', '▉', ' ', ' ', ' ', '▉', ' ', ' '},
		{'▉', ' ', ' ', ' ', '▉', ' ', '▉', '▉', '▉', ' ', ' ', ' ', '▉', ' ', ' '},
		{' ', '▉', '▉', '▉', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '▉', ' ', ' '},
		{' ', ' ', '▉', ' ', '▉', ' ', ' ', '▉', '▉', '▉', ' ', '▉', ' ', '▉', ' '},
		{' ', '▉', ' ', '▉', '▉', '▉', ' ', '▉', ' ', '▉', ' ', '▉', ' ', '▉', ' '},
		{' ', '▉', '▉', '▉', ' ', '▉', '▉', ' ', ' ', '▉', '▉', ' ', '▉', '▉', ' '},
	}
	x := 0
	fmt.Println(x)

	for h := 0; h < 10; h++ {
		for v := 0; v < 15; v++ {
			fmt.Print(string(dog[h][v]))
		}
		fmt.Println()
	}

	// radost := ""
	// chistota := ""
	// sytost := ""
	// fmt.Println(fmt.Sprintf("Radost %s\n", radost))
	// 	fmt.Println(fmt.Sprintf("Sytost %s\n", sytost))
	// 	fmt.Println(fmt.Sprintf("Chistota %s\n", chistota))

	// for {
	// 	if len(radost) == 10 {
	// 		fmt.Println("Gane over")
	// 		break
	// 	}

	// 	fmt.Println("ANIMATION")

	// 	radost+="+"
	// 	sytost+="+"
	// 	chistota+="+"
	// 	fmt.Println(fmt.Sprintf("Radost %s\n", radost))
	// 	fmt.Println(fmt.Sprintf("Sytost %s\n", sytost))
	// 	fmt.Println(fmt.Sprintf("Chistota %s\n", chistota))
		
	// 	time.Sleep(time.Second)
		


	}