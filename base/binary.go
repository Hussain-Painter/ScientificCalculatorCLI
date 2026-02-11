package base

import (
	"fmt"
	"regexp"
	"scientific/calculator/cli/renderer"
	"strconv"
)

func Binary() {
	var ch int
	for {
		err := renderer.Render("BINARY CONVERSION CALCULATOR", []string{"Binary->Octal", "Binary->Decimal", "Binary->Hexadecimal"})
		if err != nil {
			fmt.Printf("Something went terribly wrong in rendering: %e\n", err)
			return
		}
		fmt.Print("\t\t\t\t\t*****Enter your choice*****\n\t\t\t\t\t\t")

		fmt.Scanf("%d\n", &ch)
		var op string
		switch ch {
		case 1:
			fmt.Printf("\t\t\t\tEnter Operand: \n\t\t\t\t")
			fmt.Scanf("%s\n", &op)
			re, err := regexp.Compile(`^[01]+$`)

			if !re.MatchString(op) || err != nil {
				fmt.Printf("Invalid Input %s!\n", op)
				if err != nil {
					fmt.Printf("Error: %e\n", err)
				}
				ch = 4
				break
			}
			decimalVal, err := strconv.ParseInt(op, 2, 64)
			if err != nil {
				panic(fmt.Sprintf("Error while parsing number %s: %e\n", op, err))
			}
			octalStr := strconv.FormatInt(decimalVal, 8)
			fmt.Printf("Result: %s\n", octalStr)

		case 2:
			fmt.Println("OCTAL MENU")
		case 3:
			fmt.Println("DECIMAL MENU")
		}
		if ch == 4 {
			break
		}
		fmt.Printf("\t\t\t\tDo you want to continue in binary menu? (Y/N): \n\t\t\t\t")
		var cont rune
		fmt.Scanf("%c\n", &cont)
		if cont == 'N' || cont == 'n' {
			break
		}
	}
}
