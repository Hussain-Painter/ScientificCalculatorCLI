package base

import (
	"fmt"
	"regexp"
	"scientific/calculator/cli/renderer"
	"strconv"
)

func Octal() {
	var ch int
	for {
		err := renderer.Render("OCTAL CONVERSION CALCULATOR", []string{"Octal->Binary", "Octal->Decimal", "Octal->Hexadecimal"})
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
			binaryStr, err := oct2binary(op)
			if err != nil {
				ch = 4
				break
			}
			fmt.Printf("Result: %s\n", binaryStr)

		case 2:
			fmt.Printf("\t\t\t\tEnter Operand: \n\t\t\t\t")
			fmt.Scanf("%s\n", &op)
			decimalValStr, err := oct2dec(op)
			if err != nil {
				ch = 4
				break
			}
			fmt.Printf("Result: %s\n", decimalValStr)
		case 3:
			fmt.Printf("\t\t\t\tEnter Operand: \n\t\t\t\t")
			fmt.Scanf("%s\n", &op)
			hexaDecimalValStr, err := oct2hex(op)
			if err != nil {
				ch = 4
				break
			}
			fmt.Printf("Result: %s\n", hexaDecimalValStr)
		}
		if ch == 4 {
			break
		}
		fmt.Printf("\t\t\t\tDo you want to continue in octal menu? (Y/N): \n\t\t\t\t")
		var cont rune
		fmt.Scanf("%c\n", &cont)
		if cont == 'N' || cont == 'n' {
			break
		}
	}
}

func oct2binary(input string) (string, error) {
	re, err := regexp.Compile(`^[0-7]+$`)

	if !re.MatchString(input) || err != nil {
		fmt.Printf("Invalid Input %s!\n", input)
		if err != nil {
			fmt.Printf("Error: %e\n", err)
		}
		return "", err
	}
	decimalVal, err := strconv.ParseInt(input, 8, 64)
	if err != nil {
		fmt.Printf("Error while parsing number %s: %e\n", input, err)
		return "", err
	}
	binaryStr := strconv.FormatInt(decimalVal, 2)
	return binaryStr, nil

}

func oct2dec(input string) (string, error) {
	re, err := regexp.Compile(`^[0-7]+$`)

	if !re.MatchString(input) || err != nil {
		fmt.Printf("Invalid Input %s!\n", input)
		if err != nil {
			fmt.Printf("Error: %e\n", err)
		}
		return "", err
	}
	decimalVal, err := strconv.ParseInt(input, 8, 64)
	if err != nil {
		fmt.Printf("Error while parsing number %s: %e\n", input, err)
		return "", err
	}
	binaryStr := strconv.FormatInt(decimalVal, 10)
	return binaryStr, nil

}

func oct2hex(input string) (string, error) {
	re, err := regexp.Compile(`^[0-7]+$`)

	if !re.MatchString(input) || err != nil {
		fmt.Printf("Invalid Input %s!\n", input)
		if err != nil {
			fmt.Printf("Error: %e\n", err)
		}
		return "", err
	}
	decimalVal, err := strconv.ParseInt(input, 8, 64)
	if err != nil {
		fmt.Printf("Error while parsing number %s: %e\n", input, err)
		return "", err
	}
	binaryStr := strconv.FormatInt(decimalVal, 16)
	return binaryStr, nil

}
