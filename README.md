# rolldice [![travis](https://travis-ci.org/tummychow/rolldice.png)](https://travis-ci.org/tummychow/rolldice)

A simple program, in Go, that rolls some dice. Yes, I know it's trivial. I wrote it to get some practice with the language.

As of v1.1.0+, rolldice is based on [cli.go](http://github.com/codegangsta/cli). It's a pretty lightweight solution for these kinds of programs. Probably the only thing it's missing (in the context of this application) is typed arguments. Maybe I'll make a PR for it later.

---

## Usage
```
rolldice [--seed|-s <seed>] <num> <faces> [modifier]
```
Rolls `num` dice, where each die has `faces` (numbered starting at 1). If `modifier` is not given, all the dice will be printed, one roll per line. If `modifier` *is* given, then only the sum of the dice will be printed, plus the modifier.

Run the program without any arguments, or with the `-h` flag, to see help.

### Fixed seeds
If the `-s` flag is given, then `seed` will be used as the seed for [the Go PRNG](http://golang.org/pkg/rand). You can use this flag to get repeated, predictable results. The seed must be a positive integer. If it's not specified, rolldice falls back on the default behavior, which uses the system time (therefore results will differ between uses, as expected).

I don't know if the PRNG implementation is platform-specific, so maybe the output will differ from platform to platform. However, if two invocations with the same seed on the same platform give different results, that is a bug. Please open an issue.

### d-style input
Instead of specifying `num`, `faces` and `modifier` as numbers on the command line, you can specify them all in one string, using a notation that is common in many games. An example is worth a hundred words:
```
$ rolldice 3d6+0
8
$ rolldice 3d6
2
5
4
```
The invocation `rolldice 3d6+0` is equivalent to the invocation `rolldice 3 6 +0`. Similarly, `rolldice 3d6` is equivalent to `rolldice 3 6`. The `d` is not case sensitive, so `3d6` and `3D6` are both acceptable. Negative modifiers can obviously be specified with a minus sign. `rolldice 2d4-2` is equivalent to `rolldice 2 4 -2`.

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
```
# Reproducible rolls with -s
$ rolldice -s 1 2 6
6
4
$ rolldice -s 1 2 6
6
4
```
```
# Unique rolls with -u
$ rolldice -u 3 3
1
3
2
$ rolldice -u 4 3
<num> must be <= <faces> when using unique rolls: 4 3
```

---

## Todo
- Add a flag to print the individual rolls when the modifier is specified, so you can see both the dice and their sum. Great if you're too lazy to add numbers, like yours truly!

## Contributing
Pull requests welcome! Making trivial programs is a good way to get started with a language, hacking on trivial programs is another good way. Issues are also welcome, if you find a bug but aren't familiar with the Go language.

## License
MIT, see [LICENSE.md](LICENSE.md).
