# go-guessing-game

[![Go Report Card](https://goreportcard.com/badge/github.com/svetlana-rezvaya/go-guessing-game)](https://goreportcard.com/report/github.com/svetlana-rezvaya/go-guessing-game)

The app implementing the [guessing game](https://www.rosettacode.org/wiki/Guess_the_number).

## Installation

```
$ go get github.com/svetlana-rezvaya/go-guessing-game
```

## Usage

```
$ go-guessing-game -h | -help | --help
$ go-guessing-game [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-min INTEGER` &mdash; minimal secret (default: 1);
- `-max INTEGER` &mdash; maximal secret (default: 10).

## Output Example

```
Enter a number: 0
2020/10/27 19:30:00 number is too small
Enter a number: 1
2020/10/27 19:30:00 you guessed wrong
Enter a number: 2
2020/10/27 19:30:00 you guessed wrong
Enter a number: 3
You guessed right!
```

## License

The MIT License (MIT)

Copyright &copy; 2020 svetlana-rezvaya
