===========================================

Authors notes:

to run the program run 
`go run main.go`s
and follow the terminal instructions 

for testing you can use some tricks. 17 moves in the cli is a lot when you're testing stuff to in models .go you can update the base health to speed things up or manually update h.npchealth/playerhealth in helpers or main if you only want one switched. 

i tested with all the cases files like off_grid and crossed_boats but the bulk of my testing was modeled off of full_game. 

i created an additional instruction set that would target all of the spots setup by full_game. 

copy and paste to set the ship locations

`Destroyer right A1
Carrier down B2
Battleship down G4
Submarine right E6
Cruiser right J8
 ` 
 (the space at the end lets all ships be made in one batch but creates a buffer. type shipmap and it'll all be synced).

then you can spam 
`fire
g8` 
(or whatever you want)
to lose fast and see the loss condition. 

you can change the file you're using by updating the models.go useFile value

================================================

# Battleship 
The classic, grid-based game of naval warfare.

Each player starts with an empty 10x10 grid. Rows are identifed with the letters A-J, columns by the numbers 1-10. Each player places all of their ships on their grid. They then take turns firing at positions on the grid until either _all_ of their opponent's ships are sunk or their own ships are all sunk.

https://en.wikipedia.org/wiki/Battleship_(game)

## This project

Instead of a two-player version of the game this project is instead to implement a single player CLI version that accepts input files and produces the expected output as described below.

### Board

```text
  | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10|
A |___|___|___|___|___|___|___|___|___|___|
B |___|___|___|___|___|___|___|___|___|___|
C |___|___|___|___|___|___|___|___|___|___|
D |___|___|___|___|___|___|___|___|___|___|
E |___|___|___|___|___|___|___|___|___|___|
F |___|___|___|___|___|___|___|___|___|___|
G |___|___|___|___|___|___|___|___|___|___|
H |___|___|___|___|___|___|___|___|___|___|
I |___|___|___|___|___|___|___|___|___|___|
J |___|___|___|___|___|___|___|___|___|___|
```

### Ships

| Class of ship | Size |
| --- | --- |
| Carrier | 5 |
| Battleship | 4 |
| Cruiser | 3 |
| Submarine | 3 |
| Destroyer | 2 |


## API

### Types and Values:

    Direction: down, right
    Position: ex. A10, B5, F9
    Ships: Carrier, Battleship, Cruiser, Submarine, Destroyer

### Commands:
```
PLACE_SHIP
FIRE
```

#### `PLACE_SHIP`
```
  Inputs:
    Ship
    Direction
    Position
  Output:
    "Placed {Ship}"
```

#### `FIRE`
```
  Input:
    Position
  Output:
    “Hit” |
    “Miss” |
    “You sunk my {Ship}!” |
    “Game Over”
```

### Example Input Sequence

```
PLACE_SHIP Destroyer right A1
PLACE_SHIP Carrier down B2
PLACE_SHIP Battleship right J4
PLACE_SHIP Submarine right E6
PLACE_SHIP Cruiser right H1

FIRE A4
FIRE E6
FIRE E7
FIRE E8
FIRE B1
...
```

### Example Output
```
Placed Destroyer
Placed Carrier
Placed Battleship
Placed Submarine
Placed Cruiser
Miss
Hit
Hit
Hit
You sunk my Submarine!
Miss
...
You sunk my Cruiser!
...
Hit
Hit
Hit
You sunk my Destroyer!
...
Game Over
```
