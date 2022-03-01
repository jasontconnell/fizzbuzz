package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type stageRule struct {
	mod    int
	s      string
	format bool
}

func main() {
	start := time.Now()
	stg := getRules(getEstHour())

	rnglow, rnghigh := parseArgs(os.Args)

	out := runStage(rnglow, rnghigh, stg)
	for _, s := range out {
		fmt.Println(s)
	}
	fmt.Println("finished", time.Since(start))
}

func parseArgs(args []string) (int, int) {
	defStart, defEnd := 1, 100
	if len(args) == 3 {
		defStart, _ = strconv.Atoi(args[1])
		defEnd, _ = strconv.Atoi(args[2])

		if defStart <= 0 {
			defStart = 1
		}

		if defEnd > 100 {
			defEnd = 100
		}
	}

	return defStart, defEnd
}

func getEstHour() int {
	t := time.Now()
	loc, err := time.LoadLocation("EST")
	if err != nil || t.Location().String() == "EST" {
		return t.Hour()
	}

	nt := t.In(loc)
	return nt.Hour()
}

func getRules(h int) []stageRule {
	var rules []stageRule
	if h%24 < 12 {
		rules = []stageRule{
			{mod: 4, s: "Plong", format: false},
			{mod: 2, s: "Pling", format: false},
			{mod: 1, s: "%d", format: true},
		}
	} else {
		rules = []stageRule{
			{mod: 15, s: "FizzBuzz", format: false},
			{mod: 5, s: "Buzz", format: false},
			{mod: 3, s: "Fizz", format: false},
			{mod: 1, s: "%d", format: true},
		}
	}

	return rules
}

func runStage(start, end int, rules []stageRule) []string {
	output := []string{}
	for i := start; i <= end; i++ {
		for _, rule := range rules {
			if i%rule.mod == 0 {
				val := rule.s
				if rule.format {
					val = fmt.Sprintf(val, i)
				}
				output = append(output, fmt.Sprintf("%d == '%s'", i, val))
				break
			}
		}
	}
	return output
}
