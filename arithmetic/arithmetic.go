package arithmetic

import (
	"fmt"
	"math"
	"scientific/calculator/cli/renderer"
)

func Arithmetic() {
	// exec.Command("cls")
	//exec.Command("cls")
	var ch int
	for {
		err := renderer.Render("ARITHMETIC", []string{"Addition", "Subtraction", "Division", "Multiplication", "Sine", "Cos", "Tan", "Cosec", "Sec", "Cot", "Log", "Power"})
		if err != nil {
			fmt.Printf("Something has gone terribily wrong %e", err)
			return
		}
		fmt.Print("\t\t\t\t\t*****Enter your choice*****\n\t\t\t\t\t\t")

		fmt.Scanf("%d\n", &ch)

		var op1, op2 float64
		switch ch {
		case 1:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tEnter 2nd Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op2)
			fmt.Printf("\t\t\t\tResult: %g\n", op1+op2)
		case 2:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tEnter 2nd Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op2)
			fmt.Printf("\t\t\t\tResult: %g\n", op1-op2)
		case 3:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tEnter 2nd Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op2)
			fmt.Printf("\t\t\t\tResult: %g\n", op1/op2)
		case 4:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tEnter 2nd Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op2)
			fmt.Printf("\t\t\t\tResult: %g\n", op1*op2)
		case 5:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tResult: %g\n", math.Sin(op1))
		case 6:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tResult: %g\n", math.Cos(op1))
		case 7:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tResult: %g\n", math.Tan(op1))
		case 8:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tResult: %g\n", 1/math.Sin(op1))
		case 9:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tResult: %g\n", 1/math.Cos(op1))
		case 10:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tResult: %g\n", 1/math.Tan(op1))
		case 11:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tResult: %g\n", math.Log(op1))
		case 12:
			fmt.Printf("\t\t\t\tEnter 1st Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op1)
			fmt.Printf("\t\t\t\tEnter 2nd Operand: \n\t\t\t\t")
			fmt.Scanf("%g\n", &op2)
			fmt.Printf("\t\t\t\tResult: %g\n", math.Pow(op1, op2))
		}
		if ch == 13 {
			break
		}
		fmt.Printf("\t\t\t\tDo you want to continue in math menu? (Y/N): \n\t\t\t\t")
		var cont rune
		fmt.Scanf("%c\n", &cont)
		if cont == 'N' || cont == 'n' {
			break
		}
	}
}
