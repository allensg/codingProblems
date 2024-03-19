package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type (
	// Helpers will act as an entry point to anything i need for
	// data generation and test suite control
	Helpers struct {
		Scanner             *bufio.Scanner
		PlayerBoard         [][]string
		PlayerMoves         []*Move         // records fires for score and map generation
		PlayerShipLocations map[string]bool // records placements
		PlayerHealth        int
		NPCBoard            [][]string
		NPCMovesToMake      []*Move // array of moves read from the file during initalization
		NPCMovesMade        []*Move // populatd during the actual game loop used for making hit/miss maps
		NPCShipLocations    map[string]bool
		NPCHealth           int
	}
)

func (h *Helpers) InitializeNewGame() {
	// print the tutorial options
	h.printTutorial()
	// create a new boards for player and npc
	h.PopulateNPCBoard()
	h.PopulatePlayerBoard()
}

func (h *Helpers) printTutorial() {
	fmt.Print("Welcome to Battleship, would you like to see the tutorial? 'yes' or 'no': ")
	// reads user input until \n by default
	for {
		h.Scanner.Scan()
		// Holds the string that was scanned
		input := h.Scanner.Text()
		switch input {
		case "yes":
			fmt.Println(TutorialText)
			return
		case "no":
			fmt.Println("Sounds good.\n")
			return
		default:
			fmt.Println("Please answer either 'yes' or 'no'")
		}
	}
}

func (h *Helpers) PopulatePlayerBoard() {

	populatePlayerBoardPreamble()
	h.PlayerShipLocations = PlayerShipLocations
	shipPlacements := map[string]*Placement{}
	h.PlayerMoves = []*Move{}

	// initialize empty board
	playerBoard := EmptyPlayerBoard
	h.PlayerBoard = playerBoard
	h.PlayerHealth = baseHealth

	scan := bufio.NewScanner(os.Stdin)
	h.Scanner = scan
	fmt.Print("Please enter ship placement: ")
	// terminate the loop when all 5 ships have a placement
	for 5 > len(shipPlacements) {

		h.Scanner.Scan()
		// Holds the string that was scanned
		input := h.Scanner.Text()
		// ease of use case
		if strings.Trim(input, " ") == "show me" {
			getRemainingShipsToPlace(shipPlacements)
			continue
		}
		// validate and extract the elements from the input request
		placement, err := validatePlacementInput(input)
		if err != "" {
			fmt.Println(err)
			continue
		}

		// now that we have valid inputs we need to check to make sure the attempted placement is valid
		// and then update the board with the ship positions
		valid, err := h.checkPlacementInput(placement, Player)
		if valid {
			shipPlacements[placement.Ship] = placement
			playerSuccessPlacementPrint(placement, len(shipPlacements))
		} else {
			fmt.Println(err)
			continue
		}
	}

	// by this point each ship should be placed and the board should be updated with their locations
	fmt.Println("Excellent! Board Successfully Populated\n")
}

// this board will be populated from the input provided input file
func (h *Helpers) PopulateNPCBoard() {
	// initialize the board to 0
	npcBoard, npcMoves := EmptyNpcBoard, []*Move{}

	h.NPCBoard = npcBoard
	h.NPCShipLocations = NPCShipLocations
	h.NPCHealth = baseHealth

	// open the input file game
	file, err := os.Open(useFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// while we're building out the board we'll go ahead and capture the moves
	// for the rest of the game so we don't have to scan the file twice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		if row != "" {
			rowElements := strings.Split(row, " ")
			if rowElements[0] == "PLACE_SHIP" {

				position, err := boardPosToField(rowElements[3])
				if position == nil {
					log.Fatal(err)
				}

				place := &Placement{
					Ship:      rowElements[1],
					Direction: rowElements[2],
					Position:  position,
				}

				// prettyPrintPlacement(place)
				_, err = h.checkPlacementInput(place, NPC)
				if err != "" {
					log.Fatal(err)
					return
				}
			}
			if rowElements[0] == "FIRE" {
				pos, err := boardPosToField(rowElements[1])
				if err != "" {
					log.Fatal(err)
					return
				}

				move := &Move{
					Fire: &Fire{
						Position: pos,
					},
				}
				npcMoves = append(npcMoves, move)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	h.NPCMovesToMake = npcMoves

	// uncomment to see the built npc board on initalization
	// fmt.Println("firing instructions")
	// fmt.Println(len(h.NPCMoves))
	// h.PrintBoard(h.NPCBoard)

}

// checks to see if the attempted placement is valid
// splitting by space should have it as
// 0 - Ship, 1 - direction, 2 - position
func validatePlacementInput(input string) (*Placement, string) {
	inputFields := strings.Split(input, " ")
	if len(inputFields) != 3 {
		return nil, FormatErrorStr
	}

	// check the ship
	if _, found := ShipDefinitions[inputFields[0]]; !found {
		return nil, "Invalid ship entered"
	}

	// check the orientation
	if _, found := DirectionDefinitions[inputFields[1]]; !found {
		return nil, "Invalid orientation entered"
	}

	// split the origin into letter and number. check both
	location := inputFields[2]
	position, err := boardPosToField(location)
	if position == nil {
		return nil, err
	}

	placement := &Placement{
		Ship:      inputFields[0],
		Direction: inputFields[1],
		Position:  position,
	}
	return placement, ""
}

// turns a board position like 'A4' into a Position Object
func boardPosToField(input string) (*Position, string) {
	// is it populated and less than 4 letters
	if len(input) == 0 || len(input) > 3 {
		return nil, "Invalid location entered"
	}

	LetterIndex, foundLetter := BoardLetterDefinitions[string(input[0])]
	numberRemainder := string(input[1:])
	IntIndex, foundNumber := BoardNumberDefinitions[numberRemainder]
	if !foundLetter || !foundNumber {
		return nil, "location not found in hash"
	}
	position := &Position{
		FullValue:   strings.Trim(input, " "),
		LetterIndex: LetterIndex,
		IntIndex:    IntIndex,
	}

	return position, ""
}

// take in the players fire command
func (h *Helpers) Fire(input string) {

	// make this a loop in case the player enters
	// an invalid input and we need to ask again.
	// hit or miss, if the input is valid it should only be a single loop

	for {
		fmt.Println("Location: ")
		h.Scanner.Scan()
		input := h.Scanner.Text()
		in := strings.ToUpper(input)
		// ask the player for their target location
		position, err := boardPosToField(in)
		// if bad ask for it again
		if err != "" {
			continue
		}

		// log hit/miss
		status, found := h.NPCShipLocations[position.FullValue]
		if found { // hit
			// if status is true, that position has already been hit.
			// Inform player and return to main loop
			if status {
				fmt.Println("You have already hit this spot, please pick another location")
				return
			} else {
				// valid hit
				fmt.Printf("%s: Hit \n", position.FullValue)
				// mark location hit
				h.NPCShipLocations[position.FullValue] = true
				h.NPCHealth--
			}
		} else { // player missed their shot
			fmt.Printf("%s: Miss \n", position.FullValue)
		}

		// add move to list
		move := &Move{
			Fire: &Fire{
				Position: position,
			},
		}
		h.PlayerMoves = append(h.PlayerMoves, move)

		// play the next move in the npc sequence.
		h.EnemyTurn()

		return
	}

}

func (h *Helpers) EnemyTurn() {
	move := &Move{}
	// figure out what the next move to make is
	// how many moves have we made so far
	moveToMakeIndex := len(h.NPCMovesMade)
	move = h.NPCMovesToMake[moveToMakeIndex]

	position := move.Fire.Position
	status, found := h.PlayerShipLocations[position.FullValue]
	if found { // hit
		// if status is true, that position has already been hit.
		if status {
			// if this happens its an error in the instruction set it will count as a
			// missed move skip to the bottom, add it to the npc moves made arr and keep goin
			fmt.Printf("Enemy Hit: %s no damage \n", position.FullValue)
		} else {
			// valid hit
			fmt.Printf("Enemy Hit: %s \n", position.FullValue)
			// mark location hit
			h.PlayerShipLocations[position.FullValue] = true
			h.PlayerHealth--
		}
	} else { //npc missed their shot
		fmt.Printf("Enemy missed at: %s \n", position.FullValue)
	}

	// add move to list
	h.NPCMovesMade = append(h.NPCMovesMade, move)
}

// TODO
func (h *Helpers) PrintHitMap() {
	fmt.Println()
}

// takes in a placement an origin and the current board and
// attempts to replace the range of the board defined by placement with the ship defined by placement
// either update the board or returns an error if collision / offgrid
func (h *Helpers) checkPlacementInput(placement *Placement, origin string) (bool, string) {
	// take in the ship, orientation and location
	var err string
	switch origin {
	case NPC:
		shipLocations := []string{}
		h.NPCBoard, shipLocations, err = h.replaceElements(h.NPCBoard, placement)
		if err != "" {
			return false, err
		}
		h.NPCShipLocations = populateLocations(shipLocations, h.NPCShipLocations)

	case Player:
		shipLocations := []string{}
		h.PlayerBoard, shipLocations, err = h.replaceElements(h.PlayerBoard, placement)
		if err != "" {
			return false, err
		}

		h.PlayerShipLocations = populateLocations(shipLocations, h.PlayerShipLocations)
	default:
		return false, "invalid origin"
	}

	return true, ""
}

func (h *Helpers) replaceElements(board [][]string, placement *Placement) ([][]string, []string, string) {
	// make a backup copy of the board to rollback to in case we hit a collision
	var boardBeforeReplace = [][]string{}
	copy(boardBeforeReplace, board)
	// fmt.Println("replacement board")
	// h.PrintBoard(board)

	// keep track of the number of positions we've set so far
	current := 0
	shipLocations := []string{}

	// instantiate helper variables
	count := ShipDefinitions[placement.Ship]
	newMapValue := ShipMapRepresentation[placement.Ship]
	// replaceHVDiagPrint(placement, current, count, newMapValue)

	// replace the 0 values at the indexes on the board with the corresponding
	// ship number to show its been placed
	for current != count {
		x, y := getXY(placement, current)
		// check off grid
		if x > 9 || x < 0 || y > 9 || y < 0 {
			return boardBeforeReplace, shipLocations, "Error: that position would place the ship off the board. Please choose again"
		}

		toReplace := board[x][y]
		// check crossed boats
		if strings.Trim(toReplace, " ") != "0" {
			return boardBeforeReplace, shipLocations, "Error: that position is overlapping another piece. Please choose again"
		} else {
			// else replace
			board[x][y] = newMapValue
			current++
			// rebuild the input style string 'A1' from the xy coords 0,0
			letter, number := LetterBoardDefinitions[x], NumberBoardDefinitions[y]
			joined := letter + number
			shipLocations = append(shipLocations, joined)
		}
	}

	return board, shipLocations, ""
}

// determine the coordinates for the current item to be replaced based on the direction.
func getXY(placement *Placement, current int) (int, int) {
	switch placement.Direction {
	case Right:
		return placement.Position.LetterIndex, placement.Position.IntIndex + current
	case Down:
		return placement.Position.LetterIndex + current, placement.Position.IntIndex
	default:
		return -1, -1
	}
}

func (h *Helpers) PrintBoard(board [][]string) {
	fmt.Println(BoardHeader)
	fmt.Println(BoardUnderScoreBoundary)
	for column, row := range board {
		columnLetter := indexToLetterPrint[column]
		rowStr := strings.Join(row, " | ")
		fmt.Printf("%s | %s | \n", columnLetter, rowStr)
	}
	fmt.Println(BoardUnderScoreBoundary)
}
