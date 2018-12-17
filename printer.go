package bits

import (
	"fmt"

	"github.com/fatih/color"
)

// PrintMatrix prints a matrix of bytes in an easy to view
// checkerboard format
func PrintMatrix(matrix [][]byte) {
	printHeader(matrix)
	fmt.Println()
	printRows(matrix)
}

func printRows(matrix [][]byte) {
	for n, row := range matrix {
		printRow(n, row)
		printStats(n)
		fmt.Println()
	}
}

func printRow(n int, row []byte) {
	for i := 0; i < len(row); i++ {
		bt := fmt.Sprintf("%08b", row[i])
		for _, char := range bt {
			s := string(char)
			if s == "0" {
				color.Set(color.BgWhite)
				color.Set(color.FgBlack)
			} else {
				color.Set(color.FgWhite)
			}
			fmt.Printf(" %v ", string(char))
			color.Unset()
		}
	}
}

func printHeader(matrix [][]byte) {
	fmt.Print(" ")
	for n := 0; n < len(matrix[0])*8; n++ {
		fmt.Printf("%-3v", n)
	}
	fmt.Println()
}

func printStats(i int) {
	fmt.Printf("  %-3v", i)
}
