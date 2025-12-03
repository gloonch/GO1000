package main

import "fmt"

func main() {
	students := [...][3]float64{
		{4, 4, 4},
		{1, 2, 4},
		{9, 8, 9},
	}
	var sum float64

	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}

	const N = float64(len(students) + len(students[0]))
	fmt.Printf("Average Grade: %g\n", sum/N)
}
