package main

import (
	"fmt"

	"math/rand"
	"time"
	"github.com//RPG-Go/provider"
)


func main() {
	monsters()
}

// }

// func startGame() {
// 	choicePerson()

// }

// func choicePerson()  {

// 	var choice string
// 	fmt.Println("Choose your character")
// 	fmt.Println("Warrior")
// 	fmt.Println("Mage")
// 	fmt.Scan(&choice)
// 	if choice == "warrior" {
// 		fmt.Print("\033[H\033[2J")
// 		fmt.Println("You are warrior")
// 		fmt.Println("Life: ", warriorLife)
// 	} else if choice == "mage" {
// 		fmt.Print("\033[H\033[2J")
// 		fmt.Println("You are Mage")
// 		fmt.Println("Life: ", mageLife)
// 	} else {
// 		fmt.Print("\033[H\033[2J")
// 		fmt.Println("Invalid option\f")
// 		choicePerson()
// 	}
// 	actionsPerson(choice)

// }

// func actionsPerson(choice string) {
// 	var action int
// 	fmt.Println("Choose your action")
// 	fmt.Println("1 - enter the dungeon")
// 	fmt.Println("2 - my informations")
// 	fmt.Scan(&action)
// 	if action == 1 {
// 		fmt.Print("\033[H\033[2J")
// 		startDungeon()
// 	} else if action == 2 {
// 		fmt.Print("\033[H\033[2J")
// 		fmt.Println("You are a:", choice)
// 	} else {
// 		fmt.Print("\033[H\033[2J")
// 		fmt.Println("Invalid option\f")
// 	}

// }
// func actionsBattle() {
// 	var choice int
// 	fmt.Println("Choose your action")
// 	fmt.Println("1 - attack")
// 	fmt.Println("2 - run")
// 	fmt.Println("3 - use potion (cure)")
// 	fmt.Println("4 - exit dungeon")
// 	fmt.Scan(&choice)
// 	if choice == 1 {
// 		fmt.Print("\033[H\033[2J")
// 		fmt.Println("You attack the monster")

// 	} else if choice == 2 {
// 		fmt.Print("\033[H\033[2J")
// 		fmt.Println("You run away for the dungeon")
// 		monsters()
// 	} else if choice == 3 {
// 		fmt.Print("\033[H\033[2J")
// 		fmt.Println("You use a potion")
// 	} else if choice == 4 {
// 		fmt.Print("\033[H\033[2J")
// 		//actionsPerson(choicePerson())
// 	} else {
// 		fmt.Print("\033[H\033[2J")
// 		fmt.Println("Invalid option\f")
// 		actionsBattle()
// 	}
// }

// func startDungeon() {

// 	fmt.Println("You enter the dungeon")
// 	monsters()

// }

func monsters() {
	monster := []string{"goblin", "orc", "trol", "dragon"}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(provider.monsters))
	fmt.Println("you found one: ", monster[randomIndex])
	//actionsBattle()

}
