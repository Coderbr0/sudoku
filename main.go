package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

// Sudoko grid is represented as a slice of 81 bytes
// 0 means number is unknow
// grid is a global variable, so it is acceisble by main and all functions

var grid [81]byte

func main() {
	var gridpos byte // variable to hold the current position (0~80) in the grid

	args := os.Args[1:]
	numarg := len(args)
	if numarg != 9 {
		fmt.Println("Error")
		return
	}

	// parse each argument (one row) and fill the grid and test if the grid is valid

	gridpos = 0
	for i := 0; i < numarg; i++ {
		srune := []rune(args[i])
		l := len(srune)
		if l != 9 {
			fmt.Println("Error")
			return
		}
		for j := 0; j < l; j++ {
			var number byte
			// code to write:
			// check each character in the row if it is between 1~9 or is .
			// if character is invalid, exit with error
			// if character is valid, put the number in the corresponding position of the grid
			// (put 0 in the grid if character is .)

			// below code still needs to check for invalid characters (other than 1~9 or .)

                        if srune[j] != '.' {
                                number = byte(srune[j] - 48)
                                grid[gridpos] = number

			}

			// after putting each number in the grid, call the Trynum function to check if the grid is valid
			if !Trynum(gridpos, number) {
				fmt.Println("Error")
				return
			}
			gridpos++
		}
	}

	// Now the grid is filled and checked
	// so we can solve the sudoku puzzle

	// write code here to solve the puzzle using the Trynum function

	// print the grid
	for gridpos := 0; gridpos < 81; gridpos++ {
		num := grid[gridpos]
		if num == 0 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(num + '0'))
		}
		if (gridpos+1)%9 == 0 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
}

// fuctions we need

func Check9(tocheck [9]byte) bool { // check a slice of 9 int for duplicates among 1-9, ignore 0
	var result bool = true // true = no duplicates,  false = has duplicates

	// code to be written

	return result
}

func Checkrow(gridpos byte) bool { // check row for duplicates - gridpos (0~80) passed as argument
	var result bool = true // true = no duplicates,  false = has duplicates
	var checknums [9]byte

	// code to fill up checknums with the numbers from the row of the griid where the element is
	// example: checknums[0] = grid[0]

	result = Check9(checknums)
	return result
}

func Checkcolumn(gridpos byte) bool { // check column for duplicates - gridpos (0~80) passed as argument
	var result bool = true // true = no duplicates,  false = has duplicates
	var checknums [9]byte

	// code to fill up checknums with the numbers from the column

	result = Check9(checknums)
	return result
}

func Checkblock(gridpos byte) bool { // check box for duplicates - gridpos (0~80)  passed as argument
	var result bool = true // true = no duplicates,  false = has duplicates
	var checknums [9]byte

	// code to fill up checknums with the numbers from the block

	result = Check9(checknums)
	return result
}

func Trynum(gridpos byte, numb byte) bool { 	// check if numb can be placed at gridpos (0~80)
	var result bool = true 			// true = number can be placed at gridpos without conflicts, otherwise false

	grid[gridpos] = numb
	result = Checkrow(gridpos) && Checkcolumn(gridpos) && Checkblock(gridpos)
	if result == false {
		grid[gridpos]=0
	}
	return result
}
