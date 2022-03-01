# Directions

## Stage 1
Your goal is to write an application that uses the FizzBuzz algorithm. The application will take a starting number and an ending number, both specified as `0 < number <= 100`.

The application will return numbers divisible by 3 as Fizz. Numbers divisible by 5 will return as Buzz. Finally, numbers divisible by 15 will return as FizzBuzz.

Examples:
```
1 == "1"
5 == "Buzz"
6 == "Fizz"
8 == "8"
15 == 'FizzBuzz'
```

## Stage 2

We need to add some new functionality to the application. It turns out that the algorithm should change when run before noon, using a new algorithm.

Again, the application will take a starting number and an ending number, both specified as `0 < number <= 100`.

When the time is before 12PM EST, the application will return a number divisible by 2 as Pling, and numbers divisible by 4 as Plong.

Examples:
```
1 == "1"
2 == "Pling"
8 == "Plong"
15 == '15'
```