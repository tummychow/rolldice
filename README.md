# rolldice

A simple program, in Go, that rolls some dice. Yes, I know it's trivial. I wrote it to get some practice with the language.

As of v1.1.0, rolldice is based on [cli.go](http://github.com/codegangsta/cli). It's a pretty lightweight solution for these kinds of programs. Probably the only thing it's missing (in the context of this application) is typed arguments. Maybe I'll make a PR for it later.

---

## Usage
```
rolldice <num> <faces> [modifier]
```
Rolls `num` dice, where each die has `faces` (numbered starting at 1). If `modifier` is not given, all the dice will be printed, one roll per line. If `modifier` *is* given, then only the sum of the dice will be printed, plus the modifier.

Run the program without any arguments, or with the `-h` flag, to see help.

### Examples
```
$ rolldice 3 6
6
1
2
```
```
$ rolldice 2 4 -2
3
```
```
# Specify a modifier of 0, if you just want to see the sum of the dice
$ rolldice 3 6 0
15
```
```
# Use sort to print dice in order
$ rolldice 3 6 | sort
1
2
4
```

---

## Todo
- Add parsing for d-style input, eg `rolldice 3d6` to roll three six-sided dice, or `rolldice 3d6+2` to roll three six-sided dice and then add 2 to their sum.
- Add a flag to set the seed, eg `rolldice 3 6 --seed 1`, to fix the PRNG to a specific seed. Makes it easier to test the program and get reproducible results.
- Add a flag to print the individual rolls when the modifier is specified, so you can see both the dice and their sum. Great if you're too lazy to add numbers, like yours truly!

## Contributing
Pull requests welcome! Making trivial programs is a good way to get started with a language, hacking on trivial programs is another good way.

## License
MIT, see [LICENSE.md](LICENSE.md).
