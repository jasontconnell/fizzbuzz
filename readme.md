# Instructions

_Prerequisites_
1. Git
2. Go
3. Optionally Docker

## Go

[Install Go](https://go.dev/dl/)

## Git

[Install Git](https://git-scm.com/downloads)

## Docker

[Install Docker](https://www.docker.com/products/docker-desktop)

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

## Docker Compose

This project includes a web frontend which will run on localhost:5678

Alternative to the previous documented build and run methods, you can do this:

`go run main.go web`

Or

```
go build
./fizzbuzz web
```

Or

```
docker-compose build
docker-compose up
```

And then hit the website in your browser at [http://localhost:5678](localhost:5678)

# Interesting Bits

In an effort to not just have a bunch of `if` statements which check the mod value, I wanted to make it easy to change if requirements changed.

For instance, if we wanted a mod 30 to print something different (e.g. FizzBuzzBazz), it would only require adding a `StageRule` to the `GetRules` method. This could further be not runtime defined and implemented in a JSON file format.

Also, I took this approach at the start because it was ingrained in me from college to never have a method do something, and then also write output. But this lead me to easily create a web frontend for this since the main method of processing isn't just writing to stdout, but returning a value.
