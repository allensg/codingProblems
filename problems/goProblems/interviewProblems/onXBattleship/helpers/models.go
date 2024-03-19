package helpers

import (
	"fmt"
	"strings"
)

// has a large set of random movements
var fullGame = "./inputs/full_game.txt"

// will target the positions in sequence listed in fullgame
var playerLoss = "./inputs/player_loss.txt"
var useFile = fullGame

const Right = "right"
const Down = "down"
const Carrier = "Carrier"
const Battleship = "Battleship"
const Cruiser = "Cruiser"
const Submarine = "Submarine"
const Destroyer = "Destroyer"
const NPC = "NPC"
const Player = "Player"

var baseHealth = 17

type Placement struct {
	Direction string
	Position  *Position
	Ship      string
}

var TutorialText = "https://en.wikipedia.org/wiki/Battleship_(game)"

// diagnostic and compressed use print statements
var prettyPrintPlacement = func(place *Placement) {
	fmt.Printf("ship: %s, direction: %s \n", place.Ship, place.Direction)
	fmt.Printf("FullValue: %s, LetterIndex: %d, IntIndex: %d \n", place.Position.FullValue, place.Position.LetterIndex, place.Position.IntIndex)

}

var replaceHVDiagPrint = func(place *Placement, current, count int, nmv string) {
	fmt.Println("==================")
	fmt.Printf("coord: %s, x: %d, y: %d, current %d, count: %d, newMapValue: %s \n", place.Position.FullValue, place.Position.LetterIndex, place.Position.IntIndex, current, count, nmv)
	prettyPrintPlacement(place)
}

var checkPlacementInputDiagPrint = func(placement *Placement, origin string) {
	fmt.Println("+++++++++++++++++")
	prettyPrintPlacement(placement)
	fmt.Printf("Origin: %s \n", origin)
}

var playerSuccessPlacementPrint = func(placement *Placement, shipsPlaced int) {
	fmt.Printf("Carrier= 5, Battleship=4, Cruiser=3, Submarine=2, Destroyer=1\n")
	fmt.Printf("Placed %s successfully, %d ships left to place\n", placement.Ship, 5-shipsPlaced)
}

var populatePlayerBoardPreamble = func() {
	fmt.Println("Before we can start you will need to place your ships. You may enter text in the following format.")
	fmt.Println("ShipName Direction StartingLocation")
	fmt.Println("Direction: down, right")
	fmt.Println("Position: A-J + 1-10. ex: A10, B5, F9")
	fmt.Println("You may choose from the following ships")
	fmt.Print("\n")
	fmt.Println("Class of Ship | Size")

	for ship, size := range ShipDefinitions {
		fmt.Printf("%s | %d\n", ship, size)
	}

	fmt.Print("\n")
	fmt.Println("for example valid inputs might be 'Destroyer right A1' or 'Carrier down B2'")
	fmt.Println("if you want to see what ships to place enter 'show me'")
}

var getRemainingShipsToPlace = func(placedShips map[string]*Placement) {
	toPrint := []string{}
	for name := range ShipDefinitions {
		_, found := placedShips[name]
		if !found {
			toPrint = append(toPrint, name)
		}
	}

	fmt.Printf("You still need to place the following ships: %s\n", strings.Join(toPrint, ", "))
}

type Position struct {
	FullValue   string
	LetterIndex int
	IntIndex    int
}

type Fire struct {
	Position *Position
}

type Move struct {
	Place *Placement
	Fire  *Fire
}

// this will also act as the ships health for each ship
var ShipDefinitions = map[string]int{
	Carrier:    5,
	Battleship: 4,
	Cruiser:    3,
	Submarine:  3,
	Destroyer:  2,
}

var ShipMapRepresentation = map[string]string{
	Carrier:    "5",
	Battleship: "4",
	Cruiser:    "3",
	Submarine:  "2",
	Destroyer:  "1",
}

var DirectionDefinitions = map[string]bool{
	Right: true,
	Down:  true,
}

// maps letters to indexes
var BoardLetterDefinitions = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
	"I": 8,
	"J": 9,
}

// maps letters to indexes
var LetterBoardDefinitions = map[int]string{
	0: "A",
	1: "B",
	2: "C",
	3: "D",
	4: "E",
	5: "F",
	6: "G",
	7: "H",
	8: "I",
	9: "J",
}

var BoardNumberDefinitions = map[string]int{
	"1":  0,
	"2":  1,
	"3":  2,
	"4":  3,
	"5":  4,
	"6":  5,
	"7":  6,
	"8":  7,
	"9":  8,
	"10": 9,
}

var NumberBoardDefinitions = map[int]string{
	0: "1",
	1: "2",
	2: "3",
	3: "4",
	4: "5",
	5: "6",
	6: "7",
	7: "8",
	8: "9",
	9: "10",
}

// maps indexes to letters
var indexToLetterPrint = map[int]string{
	0: "A",
	1: "B",
	2: "C",
	3: "D",
	4: "E",
	5: "F",
	6: "G",
	7: "H",
	8: "I",
	9: "J",
}

// add an array of locations from replaceelemtns to the player/npc map object
// default false, as no fire moves will have been made at this point
var populateLocations = func(arr []string, hash map[string]bool) map[string]bool {
	for _, k := range arr {
		hash[k] = false
	}
	return hash
}

var BoardHeader = "  | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10|"
var BoardUnderScoreBoundary = "  | _ | _ | _ | _ | _ | _ | _ | _ | _ | _ |"

var EmptyNpcBoard = [][]string{
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
}

var EmptyPlayerBoard = [][]string{
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
	{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
}

// maps that we can quickly check to see if either npc or player "A1" string hits or misses
// true or false determines if its already been hit or not
var PlayerShipLocations = make(map[string]bool)
var NPCShipLocations = make(map[string]bool)

// errors
const FormatErrorStr = "Invalid format, please check your input."
