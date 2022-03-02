package process

import (
	"fmt"
	"time"
)

// get the current hour relative to EST
func GetEstHour() int {
	t := time.Now()
	loc, err := time.LoadLocation("EST")
	if err != nil || t.Location().String() == "EST" {
		return t.Hour()
	}

	nt := t.In(loc)
	return nt.Hour()
}

// run a set of stage rules from start to end
func RunStage(start, end int, rules []StageRule) []string {
	start, end = swap(start, end) // swap if provided in wrong order
	output := []string{}
	for i := start; i <= end; i++ {
		for _, rule := range rules {
			if i%rule.Mod == 0 {
				val := rule.Value
				if rule.Format {
					val = fmt.Sprintf(val, i)
				}
				output = append(output, fmt.Sprintf("%d == '%s'", i, val))
				break
			}
		}
	}
	return output
}

func swap(low, high int) (int, int) {
	if low > high {
		low, high = high, low
	}
	return low, high
}

// get stage rules based on the provided hour
func GetRules(h int) []StageRule {
	var rules []StageRule
	if h%24 < 12 {
		rules = []StageRule{
			{Mod: 4, Value: "Plong", Format: false},
			{Mod: 2, Value: "Pling", Format: false},
			{Mod: 1, Value: "%d", Format: true},
		}
	} else {
		rules = []StageRule{
			{Mod: 15, Value: "FizzBuzz", Format: false},
			{Mod: 5, Value: "Buzz", Format: false},
			{Mod: 3, Value: "Fizz", Format: false},
			{Mod: 1, Value: "%d", Format: true},
		}
	}

	return rules
}
