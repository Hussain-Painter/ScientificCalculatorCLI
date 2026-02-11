package renderer

import (
	"errors"
	"fmt"
	"scientific/calculator/cli/constants"
	"strings"
)

func Render(header string, options []string) error {
	fmt.Printf("\t\t\t\t\t%s\n", constants.Decorator)
	fmt.Printf("\t\t\t\t\t   %s\n", header)
	fmt.Printf("\t\t\t\t\t%s\n", constants.Decorator)
	if len(options) < 1 {
		return errors.New("At least 1 option should be provided")
	}
	options = append(options, "Exit")
	var i int
	for i = 0; i < len(options)-1; i += 2 {
		padding1 := 2
		padding2 := 2
		if i >= 10 {
			padding1 = 3
		}
		if i+1 >= 10 {
			padding2 = 3
		}

		fmt.Printf("\t\t\t%s\t\t\t%s\n", strings.Repeat("_", 31), strings.Repeat("_", 31))
		fmt.Printf("\t\t\t|%s|\t\t\t|%s|\n", strings.Repeat(" ", 28), strings.Repeat(" ", 28))
		lSpaceCount1 := (28 - len(options[i]) - padding1) / 2
		rSpaceCount1 := 28 - len(options[i]) - padding1 - lSpaceCount1
		lSpaceCount2 := (28 - len(options[i+1]) - padding2) / 2
		rSpaceCount2 := 28 - len(options[i+1]) - padding2 - lSpaceCount2

		fmt.Printf("\t\t\t|%s|\t\t\t|%s|\n", fmt.Sprintf("%s%d.%s%s", strings.Repeat(" ", lSpaceCount1), i+1, options[i], strings.Repeat(" ", rSpaceCount1)), fmt.Sprintf("%s%d.%s%s", strings.Repeat(" ", lSpaceCount2), i+2, options[i+1], strings.Repeat(" ", rSpaceCount2)))
		fmt.Printf("\t\t\t|%s|\t\t\t|%s|\n", strings.Repeat("_", 28), strings.Repeat("_", 28))
	}
	fmt.Println()
	if len(options)%2 != 0 {
		padding := 2
		if i >= 10 {
			padding = 3
		}
		fmt.Printf("\t\t\t%s\n", strings.Repeat("_", 31))
		fmt.Printf("\t\t\t|%s|\n", strings.Repeat(" ", 28))
		n := len(options) - 1
		lSpaceCount1 := (28 - len(options[n]) - padding) / 2
		rSpaceCount1 := 28 - len(options[n]) - padding - lSpaceCount1
		fmt.Printf("\t\t\t|%s|\n", fmt.Sprintf("%s%d.%s%s", strings.Repeat(" ", lSpaceCount1), n+1, options[n], strings.Repeat(" ", rSpaceCount1)))
		fmt.Printf("\t\t\t|%s|\n", strings.Repeat("_", 28))

	}
	return nil
}
