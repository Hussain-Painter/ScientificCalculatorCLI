package main

import (
	"fmt"
	"scientific/calculator/cli/arithmetic"
	"scientific/calculator/cli/base"
	"scientific/calculator/cli/renderer"
)

func main() {
	// exec.Command("cls")

	var ch int

	for {
		err := renderer.Render("SCIENTIFIC CALCULATOR", []string{"Math", "Base", "Postfix Eval."})
		if err != nil {
			fmt.Printf("Something has gone terribily wrong in rendering %e", err)
			return
		}
		fmt.Print("\t\t\t\t\t*****Enter your choice*****\n\t\t\t\t\t\t")
		fmt.Scanf("%d\n", &ch)
		switch ch {
		case 1:
			arithmetic.Arithmetic()

		case 2:
			base.Base()
		case 4:
			return
		}
		fmt.Printf("\t\t\t\tDo you want to continue? (Y/N): \n\t\t\t\t")
		var cont rune
		fmt.Scanf("%c\n", &cont)
		if cont == 'N' || cont == 'n' {
			break
		}
	}
}
