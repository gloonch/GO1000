package main

import (
	"fmt"
)

func main() {
	type placeholder [5]string

	// no difference [5]string OR placeholder{
	zero := placeholder{
		"000",
		"0 0",
		"0 0",
		"0 0",
		"000",
	}

	one := placeholder{
		"00 ",
		" 0 ",
		" 0 ",
		" 0 ",
		"000",
	}

	two := placeholder{
		"000",
		"  0",
		"000",
		"0  ",
		"000",
	}

	three := placeholder{
		"000",
		"  0",
		"000",
		"  0",
		"000",
	}

	four := placeholder{
		"0 0",
		"0 0",
		"000",
		"  0",
		"  0",
	}
	five := placeholder{
		"000",
		"0  ",
		"000",
		"  0",
		"000",
	}

	six := placeholder{
		"0  ",
		"0  ",
		"000",
		"0 0",
		"000",
	}

	seven := placeholder{
		"000",
		"  0",
		"  0",
		"  0",
		"  0",
	}
	eight := placeholder{
		"000",
		"0 0",
		"000",
		"0 0",
		"000",
	}
	nine := placeholder{
		"000",
		"0 0",
		"000",
		"  0",
		"  0",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	//same as :
	//	for line := 0; line < 5; line++
	for line := range digits[0] {
		for digit := range digits {
			fmt.Print(digits[digit][line], "   ")
		}
		fmt.Println()
	}

}
