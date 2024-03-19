package main

import (
	"battleship/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// setup

	// set to true when a win condition has been met
	scan := bufio.NewScanner(os.Stdin)

	h := &helpers.Helpers{
		Scanner: scan,
	}

	// go through and setup the npc and player boards
	h.InitializeNewGame()

	// after everything is initialized alternate making player and npc moves until
	// one or the other has no ships left

	printAvailableCommands()

	for {

		// if the player or npc run out of ship health end the game
		if h.PlayerHealth == 0 {
			fmt.Println("You are out of health, you lose!")
			return
		}

		if h.NPCHealth == 0 {
			fmt.Println("NPC has no more ships, you win!")
			return
		}

		// if we are out of moves for the npc to make end the game
		if len(h.NPCMovesMade) == len(h.NPCMovesToMake) {
			fmt.Println("NPC Has no more moves to make, you win!")
			return
		}

		fmt.Print("Enter Command: ")
		// reads user input until \n by default
		h.Scanner.Scan()
		// Holds the string that was scanned
		input := h.Scanner.Text()

		switch strings.ToLower(input) {
		case "fire":
			h.Fire(input)
		case "shipmap":
			h.PrintBoard(h.PlayerBoard)
		case "hitmap":
			fmt.Println("hitmap")
		case "showhealth":
			fmt.Printf("Health remaining for player: %d | npc: %d \n", h.PlayerHealth, h.NPCHealth)
		case "options":
			printAvailableCommands()
		case "enemy":
			fmt.Println("You have no enemies...")
		case "exit":
			fmt.Println("Thank you for playing")
			return
		default:
			fmt.Println("Please answer with one of the following values")
			printAvailableCommands()
		}
	}

}

func printAvailableCommands() {
	fmt.Println("Fire - Lets you enter a firing coordinate")
	fmt.Println("ShipMap - Shows the position of your ships")
	fmt.Println("HitMap - Shows the enemy board with hits and misses - does not work")
	fmt.Println("ShowHealth - Shows the remaining health for player and npc")
	fmt.Println("Options - Shows this menu again")
	fmt.Println("exit - exit game")
}
