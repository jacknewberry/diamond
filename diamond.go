package diamond

import (
	. "strings"
)

//ToDiamond ...
// return a slice of strings that when printed form a diamond!
//i
//0    A
//1 --B B     outerPadding = 2
//2  C   C
//3 D     D
//2  C---C    innerPadding = 3
//1   B B
//0    A
//
func ToDiamond(letter string) []string {
	index := Index("ABCDEFGHIJKLMNOPQRSTUVWXYZ", letter)

	// Get the first line:
	output := []string{firstOrLastline(index)}
	if index == 0 {
		return output
	}

	// Get the first slope of the diamond (and the middle point):
	for i := 1; i <= index; i++ {
		innerPadding := (i * 2) - 1
		outerPadding := index - i
		thisLetter := letterForIndex(i)
		output = append(output, anotherLine(thisLetter, outerPadding, innerPadding))
	}

	// Get the second slope of the diamond:
	for i := index - 1; i > 0; i-- {
		innerPadding := (i * 2) - 1
		outerPadding := index - i
		thisLetter := letterForIndex(i)
		output = append(output, anotherLine(thisLetter, outerPadding, innerPadding))
	}

	// Get the last line:
	output = append(output, firstOrLastline(index))

	return output
}

func letterForIndex(index int) string {
	return string("ABCDEFGHIJKLMNOPQRSTUVWXYZ"[index])
}

func firstOrLastline(index int) string {
	return Join([]string{Repeat(" ", index), "A"}, "")
}

func anotherLine(letter string, outerPadding, innerPadding int) string {
	return Join([]string{Repeat(" ", outerPadding), letter, Repeat(" ", innerPadding), letter}, "")
}
