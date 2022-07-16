package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BankHeist() {
	rand.Seed(time.Now().UnixNano())
	var isHeistOn bool = true
	eludedGuards := rand.Intn(100)

	if eludedGuards < 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn && openedVault < 70 {
		isHeistOn = false
		fmt.Println("the vault can't be opened")
	}

	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Heist is failed")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2, 3:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open")
		default:
			fmt.Println("Start the getaway car!")
		}

		if isHeistOn {
			amtStolen := 1000 + rand.Intn(20000)
			fmt.Println("amtStolen = ", amtStolen)
		}
	}

	fmt.Println(isHeistOn)
}

func main() {
	BankHeist()
}
