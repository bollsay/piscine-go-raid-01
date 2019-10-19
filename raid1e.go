package solutions

import (
	"github.com/01-edu/z01"
)

func Raid1e(a int, b int) {
	for i := 0; i < b; i++ {
		if a > 0 {
			for j := 0; j < a; j++ {
				if (i == 0 && j == 0) || (j == a-1 && i == b-1 && j > 0 && i > 0) {
					z01.PrintRune('A')
				} else if (i == 0 && j == a-1) || (i == b-1 && j == 0) {
					z01.PrintRune('C')
				} else if i == 0 || i == b-1 || j == 0 || j == a-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
