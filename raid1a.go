package student

import (
	"github.com/01-edu/z01"
)

func Raid1a(a int, b int) {
	if a > 0 && b > 0 {
		for i := 0; i < b; i++ {
			for j := 0; j < a; j++ {
				if (i == 0 && j == 0) || (i == 0 && j == a-1) || (i == b-1 && j == 0) || (j == a-1 && i == b-1) {
					z01.PrintRune('o')
				} else if i == 0 || i == b-1 {
					z01.PrintRune('-')
				} else if j == 0 || j == a-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
