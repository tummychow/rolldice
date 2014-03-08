package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Roll rolls n dice with values in the range [1,d], and returns
// their values in the order they were rolled.
func roll(n, d int) []int {
	dice := make([]int, n)
	for i := range dice {
		dice[i] = rand.Intn(d) + 1
	}
	return dice
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 3 {
		fmt.Printf("Not enough arguments\n")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Printf("First argument must be non-negative integer\n")
		return
	}

	f, err := strconv.Atoi(os.Args[2])
	if err != nil || f <= 0 {
		fmt.Printf("Second argument must be positive integer\n")
		return
	}

	// fmt.Printf("Rolling %dd%d+%d\n", n, f, s)

	dice := roll(n, f)

	if len(os.Args) > 3 {
		s, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("Third argument must be integer\n")
			return
		}

		for i := range dice {
			s += dice[i]
		}
		fmt.Printf("%d\n", s)
	} else {
		for i := range dice {
			fmt.Printf("%d\n", dice[i])
		}
	}
}
