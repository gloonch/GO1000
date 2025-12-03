package main

import (
	"fmt"
	"time"
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

	colon := placeholder{
		"     ",
		"  0  ",
		"     ",
		"  0  ",
		"     ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	now := time.Now()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()

	// [8][5]string -> [8]placeholder
	clock := [...]placeholder{
		digits[hour/10], digits[hour%10],
		colon,
		digits[min/10], digits[min%10],
		colon,
		digits[sec/10], digits[sec%10],
	}

	//same as :
	//	for line := 0; line < 5; line++
	for line := range clock[0] {
		for digit := range clock {
			fmt.Print(clock[digit][line], "   ")
		}
		fmt.Println()
	}

}
