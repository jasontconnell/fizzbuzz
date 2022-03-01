# Instructions

_Prerequisites_
1. Git
2. Go

## Install Go

Install Go by downloading it from [go.dev](https://go.dev/dl/)

## Install Git

Install Git by downloading it from [git-scm.com](https://git-scm.com/downloads)

## Clone Repository

`git clone https://github.com/jasontconnell/fizzbuzz.git`

Change directory into fizzbuzz

`cd fizzbuzz`

## Test

Run tests with `go test -v`

## Compile & Run

Finally, compile the project and run it. It takes 2 arguments, space separated numbers for the start and end values.

`go run main.go 5 20`

Or

```
go build
./fizzbuzz 5 20
```

Will run fizzbuzz from 5 to 20, inclusive.

# Interesting Bits

In an effort to not just have a bunch of `if` statements which check the mod value, I wanted to make it easy to change if requirements changed.

For instance, if we wanted a mod 30 to print something different (e.g. FizzBuzzBazz), it would only require adding a `stageRule` to the `getRules` method. This could further be not runtime defined and implemented in a JSON file format.


