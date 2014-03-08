# rolldice

A simple program, in Go, that rolls some dice. Yes, I know it's trivial. I wrote it to get some practice with the language.

---

## Usage
```
rolldice <num> <faces> [modifier]
```
Rolls `num` dice, where each die has `faces` (numbered starting at 1). If `modifier` is not given, all the dice will be printed, one roll per line. If `modifier` *is* given, then only the sum of the dice will be printed, plus the modifier.
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
- Port to a command-line package (eg [flags](http://github.com/jessevdk/go-flags) or [cli](http://github.com/codegangsta/cli)). Right now these packages don't have good support for typed non-flag arguments, otherwise I'd already be using one.

## Contributing
Pull requests welcome! Making trivial programs is a good way to get started with a language, hacking on trivial programs is another good way.

## License
MIT, see [LICENSE.md](LICENSE.md).
