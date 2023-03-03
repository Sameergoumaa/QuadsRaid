package piscine

import "github.com/01-edu/z01"

func QuadE(x,y int){
	row := y
	column := x
	if x > 0 && y > 0 {
	for r := 1; r <= row; r++ {
		for c := 1; c <= column; c++ {
			if r == 1 && c == 1 {
				z01.PrintRune('A')		
			} else if (c == 1 || c == column || r == 1 || r == row) {
				//z01.PrintRune(r,c)
				if (r > 1 && r < row && c == 1) || (r > 1 && r < row && c == column) {
					z01.PrintRune('B')
				} else if c == column && r == 1{ 
					z01.PrintRune('C')
				} else if c == 1 && r == row {
					z01.PrintRune('C')
				} else if c == column && r == row{
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			} else {
				z01.PrintRune(' ')
			}
		}	
		z01.PrintRune('\n')
		}
	}
}
