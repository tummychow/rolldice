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
        println("  Usage:")
        println("      rolldice <num> <faces> [modifier]")
        println("  Rolls <num> dice, each with <faces> number of faces in range [1, <faces>].\n")

        println("  If [modifier] is not given, the dice are printed, one per line.")
        println("  If [modifier] is given, the sum of all the dice, plus the modifier, is")
        println("  printed. The individual rolls will not be printed.\n")

		println("  <num> must be a non-negative integer. <faces> must be a positive integer.")
        println("  [modifier] must be an integer (can be any sign, or zero).")

		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		println("<num> must be non-negative integer")
		os.Exit(1)
	}

	f, err := strconv.Atoi(os.Args[2])
	if err != nil || f <= 0 {
		println("<faces> must be positive integer")
		os.Exit(1)
	}

	dice := roll(n, f)

	if len(os.Args) > 3 {
		s, err := strconv.Atoi(os.Args[3])
		if err != nil {
			println("[modifier] must be integer")
			os.Exit(1)
		}

		for i := range dice {
			s += dice[i]
		}
		fmt.Println(s)
	} else {
		for i := range dice {
			fmt.Println(dice[i])
		}
	}
}
