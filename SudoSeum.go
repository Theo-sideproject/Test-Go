package main

import (
	"fmt"
	"os"
)

func findSquare(x int, y int) int {
	x = x
	//fmt.Println("x  :", x, " y/3 : ", int(x/3), " x :", int(x/3))
	if x < 4 {
		return int(y/3) + 1
	} else if x > 3 && x < 7 {
		return int(y/3) + 4
	} else {
		return int(y/3) + 7
	}
}

func main() {
	var sudoku [9][9]int
	if len(os.Args) == 10 {
		//fmt.Println(len(os.Args))
		for i := 1; i < len(os.Args); i++ {
			if len(os.Args[i]) == 9 {
				for j := 0; j < len(os.Args[i]); j++ {
					if (rune(os.Args[i][j]) >= 48 && rune(os.Args[i][j]) <= 57) || rune(os.Args[i][j]) == 46 {
						if rune(os.Args[i][j]) >= 48 && rune(os.Args[i][j]) <= 57 {
							sudoku[i-1][j] = int(os.Args[i][j]) - 48

							fmt.Println("s : ", findSquare(i, j), " x :", i)
						}
					} else {
						fmt.Println("Bug Arg", i, ", that length is not 9 AND you don't write something that isn't a number OR a lettre except '.'")
						return
					}
				}
			} else {
				fmt.Println("Bug Arg", i, ", that length is not 9")
			}
		}
	} else {
		fmt.Println("Error : Arg not enough")
		return
	}

	fmt.Println(sudoku[1][1])
	//
	//for i := 0; i < 9; i++ {
	//	fmt.Println(os.Args[i])
	//}

}
