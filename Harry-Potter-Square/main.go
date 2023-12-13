package main

import "fmt"

func main() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			ThirdSpell(i, j)
		}
		fmt.Println()
	}
}

func FirstSpell(i, j int) {
	if i > j {
		fmt.Print("# ")
	} else {
		fmt.Print(". ")
	}
}

func SecondSpell(i, j int) {
	if i == j {
		fmt.Print("# ")
	} else {
		fmt.Print(". ")
	}
}

func ThirdSpell(i, j int) {
	if i < j {
		fmt.Print("# ")
	} else {
		fmt.Print(". ")
	}
}
