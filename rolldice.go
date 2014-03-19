package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

// roll rolls n dice with faces in the range [1,d], and returns
// their values in the order they were rolled.
func roll(n, d int) []int {
	dice := make([]int, n)
	for i := range dice {
		dice[i] = rand.Intn(d) + 1
	}
	return dice
}

// rollUnique rolls n dice with faces in the range [1,d], such that no
// two dice have the same roll, and returns their values in the order
// they were rolled. Since the rolls must be unique, it is required that n<=d.
func rollUnique(n, d int) []int {
	// load array with all possible faces, 1,2,3,4... etc
	dice := make([]int, d)
	for i := range dice {
		dice[i] = i + 1
	}

	// durstenfeld/knuth/fisher-yates shuffle
	// we only need n numbers, so we don't need the full number of iterations
	for i := d - 1; i > d-n; i-- {
		j := rand.Intn(i + 1)
		temp := dice[i]
		dice[i] = dice[j]
		dice[j] = temp
	}

	// the numbers at the end of the array are the shuffled ones
	return dice[d-n:]
}

// dString converts a string of the form "3d6+2" or similar into three strings,
// for the three numbers therein. The regex used will also match "3d6", "3D6",
// "3D6-12", etc.
// If the string does not match the expected form, nil is returned.
func dString(d string) []string {
	matches := regexp.MustCompile(`^(\+?\d+)[dD](\+?\d*[1-9]\d*)([+-]\d+)?$`).FindStringSubmatch(d)

	if matches == nil {
		return nil
	}

	// drop the leading element, which is the full match
	matches = matches[1:]

	// if the modifier is not included, then its submatch will be empty
	// this discards it from the slice
	if matches[2] == "" {
		matches = matches[:2]
	}

	return matches
}

func main() {
	app := cli.NewApp()
	app.Name = "rolldice"
	app.Usage = "it rolls dice"
	app.Version = "1.4.0"

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

   <num>, <faces> and [modifier] can also be specified in one string, such as
   "3d6" or "2D4-2".

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
		cli.BoolFlag{"unique, u", "require unique rolls"},
	}

	app.Action = func(c *cli.Context) {
		// by default, derive input from command line args
		diceArgs := c.Args()

		// but if the command line args take the form of a d-style input,
		// then replace the command line args with the parsed form of that
		// d-string
		if len(c.Args()) == 1 {
			matches := dString(c.Args()[0])
			if matches == nil {
				fmt.Println("dice string is malformed:", c.Args()[0])
				os.Exit(1)
			}
			diceArgs = matches
		} else if len(c.Args()) < 2 {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}

		// parse num and faces from the args
		n, err := strconv.Atoi(diceArgs[0])
		if err != nil || n < 0 {
			fmt.Println("<num> must be non-negative integer:", diceArgs[0])
			os.Exit(1)
		}
		f, err := strconv.Atoi(diceArgs[1])
		if err != nil || f <= 0 {
			fmt.Println("<faces> must be positive integer:", diceArgs[1])
			os.Exit(1)
		}

		// load seed
		seed := int64(c.GlobalInt("seed"))
		if seed <= 0 {
			// default seed: current time
			seed = time.Now().UnixNano()
		}
		rand.Seed(seed)

		// generate dice
		var dice []int
		if c.GlobalBool("unique") {
			if n > f {
				fmt.Println("<num> must be <= <faces> when using unique rolls:", n, f)
				os.Exit(1)
			}
			dice = rollUnique(n, f)
		} else {
			dice = roll(n, f)
		}

		if len(diceArgs) > 2 {
			// sum-style output - first parse modifier
			s, err := strconv.Atoi(diceArgs[2])
			if err != nil {
				fmt.Println("[modifier] must be integer:", diceArgs[2])
				os.Exit(1)
			}

			for _, die := range dice {
				s += die
			}
			fmt.Println(s)
		} else {
			// die-by-die output
			for _, die := range dice {
				fmt.Println(die)
			}
		}
	}

	app.Run(os.Args)
}
