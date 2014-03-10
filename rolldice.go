package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Roll rolls n dice with faces in the range [1,d], and returns
// their values in the order they were rolled.
func roll(n, d int) []int {
	dice := make([]int, n)
	for i := range dice {
		dice[i] = rand.Intn(d) + 1
	}
	return dice
}

func main() {
	app := cli.NewApp()
	app.Name = "rolldice"
	app.Usage = "it rolls dice"
	app.Version = "1.2.0"

	// tweaked help text to be structurally similar to cli.go defaults,
	// but more informative for this application
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [global options] <num> <faces> [modifier]

   Rolls <num> dice, each with <faces> number of faces in range [1, <faces>].

   If [modifier] is not given, the dice are printed, one per line.
   If [modifier] is given, the sum of all the dice, plus the modifier, is
   printed. The individual rolls will not be printed.

   <num> must be a non-negative integer. <faces> must be a positive integer.
   [modifier] must be an integer (can be any sign, or zero).

VERSION:
   {{.Version}}

GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	app.Flags = []cli.Flag{
		cli.IntFlag{"seed, s", 0, "set the seed for the PRNG"},
	}

	app.Action = func(c *cli.Context) {
		if len(c.Args()) < 2 {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}

		n, err := strconv.Atoi(c.Args()[0])
		if err != nil || n < 0 {
			println("<num> must be non-negative integer")
			os.Exit(1)
		}

		f, err := strconv.Atoi(c.Args()[1])
		if err != nil || f <= 0 {
			println("<faces> must be positive integer")
			os.Exit(1)
		}

		seed := int64(c.GlobalInt("seed"))
		if seed <= 0 {
			// default seed: current time
			seed = time.Now().UnixNano()
		}
		rand.Seed(seed)

		dice := roll(n, f)

		if len(c.Args()) > 2 {
			s, err := strconv.Atoi(c.Args()[2])
			if err != nil {
				println("[modifier] must be integer")
				os.Exit(1)
			}

			for _, die := range dice {
				s += die
			}
			fmt.Println(s)
		} else {
			for _, die := range dice {
				fmt.Println(die)
			}
		}
	}

	app.Run(os.Args)
}
