package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

// Sudoko grid is represented as a slice of 81 bytes
// 0 means number is unknown
// grid is a global variable, so it is accesible by main and all functions

var grid [81]byte

func main() {
	var gridpos byte // variable to hold the current position (0~80) in the grid

	args := os.Args[1:]
	numarg := len(args)
	if numarg != 9 { // if not 9 args, exit with error
		fmt.Println("Error")
		return
	}

	// parse each argument (one row) and fill the grid and test if the grid is valid

	gridpos = 0
	for i := 0; i < numarg; i++ {
		srune := []rune(args[i])
		l := len(srune)
		if l != 9 { // if not 9 characters, exit with error
			fmt.Println("Error")
			return
		}
		for j := 0; j < l; j++ {
			var number byte
			// check each character in the row if it is between 1~9 or is .
			// if character is invalid, exit with error
			// if character is valid, put the number in the corresponding position of the grid
			// leave 0 in the grid if character is .)

			if (srune[j] >= '1' && srune[j] <= '9') || srune[j] == '.' {
				if srune[j] != '.' {
					number = byte(srune[j] - 48)
					grid[gridpos] = number

					// after putting each number in the grid, call the Trynum function to check if the grid is valid
					if !Trynum(gridpos, number) {
						fmt.Println("Error")
						return
					}
				}
				gridpos++
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	// Now the grid is filled and checked
	// so we can solve the sudoku puzzle

	// solve by recursively calling Trynext
	gridpos = 0
	if !Trynext(gridpos) { // if puzzle is unsolvable, exit with error
		fmt.Println("Error")
		return
	}

	// puzzle solved, print the grid
	Printgrid()
}

// fuctions we need

func Trynext(gridpos byte) bool { // try to fill the next unsolved grid position with numbers from 1~9
	//	var result bool = true // true = next position can be successfully filled or no more available unsolved positions, otherwise false

	var i, j byte

	// find next available grid poition (containing 0) from current position
	for i = gridpos; i < 81; i++ {
		if grid[i] == 0 {
			for j = 1; j <= 9; j++ { // try numbers from 1~9
				if Trynum(i, j) { // if number can be placed, try next position
					// for debugging			Printgrid()
					// to see solving progress		fmt.Println()
					result := Trynext(i) // recurse to try next grid position
					if result {
						return true
					}
				}
			}
			grid[i] = 0
			return false
		}
	}
	return true
}

func Check9(tocheck [9]byte) bool { // check a slice of 9 int for duplicates among 1-9, ignore 0
	var result bool = true // true = no duplicates,  false = has duplicates

	for i := 0; i < 9; i++ {
		if tocheck[i] != 0 {
			for j := i + 1; j < 9; j++ {
				if tocheck[j] == tocheck[i] {
					result = false
					break
				}
			}
		}
	}
	return result
}

func Checkrow(gridpos byte) bool { // check row for duplicates - gridpos (0~80) passed as argument
	var result bool = true // true = no duplicates,  false = has duplicates
	var checknums [9]byte

	// calculate grid position of left of the row
	s := (gridpos / 9) * 9

	// fill up checknums with the numbers from the row
	for i := 0; i < 9; i++ {
		checknums[i] = grid[s]
		s++
	}

	result = Check9(checknums)
	return result
}

func Checkcolumn(gridpos byte) bool { // check column for duplicates - gridpos (0~80) passed as argument
	var result bool = true // true = no duplicates,  false = has duplicates
	var checknums [9]byte

	// calculate grid position of top of the column
	s := gridpos % 9

	// fill up checknums with the numbers from the column
	for i := 0; i < 9; i++ {
		checknums[i] = grid[s]
		s = s + 9
	}

	result = Check9(checknums)
	return result
}

func Checkblock(gridpos byte) bool { // check 9x9 block for duplicates - gridpos (0~80)  passed as argument
	var result bool = true // true = no duplicates,  false = has duplicates
	var checknums [9]byte

	// calculate grid position of top left hand corner of the block
	r := (gridpos / 27) * 27     // row 0 or 3 or 6
	c := ((gridpos % 9) / 3) * 3 // column 0 or 3 or 6
	s := r + c

	//  fill up checknums with the numbers from the block
	for i := 0; i < 9; i++ {
		checknums[i] = grid[s]
		s++
		if (i+1)%3 == 0 { // skip to next row after 3 numbers
			s = s + 6
		}
	}
	result = Check9(checknums)
	return result
}

func Trynum(gridpos byte, numb byte) bool { // check if numb can be placed at gridpos (0~80)
	var result bool = true // true = number can be placed at gridpos without conflicts, otherwise false

	grid[gridpos] = numb
	result = Checkrow(gridpos) && Checkcolumn(gridpos) && Checkblock(gridpos)
	if result == false {
		grid[gridpos] = 0
	}
	return result
}

func Printgrid() {
	// print the grid
	for i := 0; i < 81; i++ {
		num := grid[i]
		if num == 0 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(num + '0'))
		}
		if (i+1)%9 == 0 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	//	z01.PrintRune('\n')
}
