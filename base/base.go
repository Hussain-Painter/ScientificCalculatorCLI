package base

import (
	"fmt"
	"scientific/calculator/cli/renderer"
)

func Base() {
	var ch int
	for {
		err := renderer.Render("BASE CONVERSION CALCULATOR", []string{"Binary", "Octal", "Decimal", "Hexadecimal"})
		if err != nil {
			fmt.Printf("Something went terribly wrong in rendering: %e\n", err)
			return
		}
		fmt.Print("\t\t\t\t\t*****Enter your choice*****\n\t\t\t\t\t\t")

		fmt.Scanf("%d\n", &ch)

		switch ch {
		case 1:
			Binary()
		case 2:
			Octal()
		case 3:
			fmt.Println("DECIMAL MENU")
		case 4:
			fmt.Println("HEXADECIMAL MENU")
		}
		if ch == 5 {
			break
		}
		fmt.Printf("\t\t\t\tDo you want to continue in base conversion menu? (Y/N): \n\t\t\t\t")
		var cont rune
		fmt.Scanf("%c\n", &cont)
		if cont == 'N' || cont == 'n' {
			break
		}
	}
}
