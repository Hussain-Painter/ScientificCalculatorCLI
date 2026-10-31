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
			octalStr, err := bin2Octal(op)
			if err != nil {
				ch = 4
				break
			}
			fmt.Printf("Result: %s\n", octalStr)

		case 2:
			fmt.Printf("\t\t\t\tEnter Operand: \n\t\t\t\t")
			fmt.Scanf("%s\n", &op)
			decimalValStr, err := bin2Dec(op)
			if err != nil {
				ch = 4
				break
			}
			fmt.Printf("Result: %s\n", decimalValStr)
		case 3:
			fmt.Printf("\t\t\t\tEnter Operand: \n\t\t\t\t")
			fmt.Scanf("%s\n", &op)
			hexaDecimalValStr, err := bin2Hex(op)
			if err != nil {
				ch = 4
				break
			}
			fmt.Printf("Result: %s\n", hexaDecimalValStr)
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

func bin2Octal(input string) (string, error) {

	re, err := regexp.Compile(`^[01]+$`)

	if !re.MatchString(input) || err != nil {
		fmt.Printf("Invalid Input %s!\n", input)
		if err != nil {
			fmt.Printf("Error: %e\n", err)
		}
		return "", err
	}
	decimalVal, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		fmt.Printf("Error while parsing number %s: %e\n", input, err)
		return "", err
	}
	octalStr := strconv.FormatInt(decimalVal, 8)
	return octalStr, nil
}

func bin2Dec(input string) (string, error) {
	re, err := regexp.Compile(`^[01]+$`)
	if !re.MatchString(input) || err != nil {
		fmt.Printf("Invalid Input %s!\n", input)
		if err != nil {
			fmt.Printf("Error: %e\n", err)
		}
		return "", err
	}
	decimalVal, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		fmt.Printf("Error while parsing number %s: %e\n", input, err)
		return "", err
	}
	return strconv.FormatInt(decimalVal, 10), nil
}

func bin2Hex(input string) (string, error) {
	re, err := regexp.Compile(`^[01]+$`)
	if !re.MatchString(input) || err != nil {
		fmt.Printf("Invalid Input %s!\n", input)
		if err != nil {
			fmt.Printf("Error: %e\n", err)
		}
		return "", err
	}
	decimalVal, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		fmt.Printf("Error while parsing number %s: %e\n", input, err)
		return "", err
	}
	return strconv.FormatInt(decimalVal, 16), nil
}
