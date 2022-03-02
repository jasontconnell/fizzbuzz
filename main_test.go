package main

import (
	"strings"
	"testing"

	"github.com/jasontconnell/fizzbuzz/process"
)

func TestRules(t *testing.T) {
	tests := []struct {
		hour     int
		expected string
	}{
		{hour: 10, expected: "Plong"},
		{hour: 12, expected: "FizzBuzz"},
		{hour: 15, expected: "FizzBuzz"},
	}

	for _, test := range tests {
		rules := process.GetRules(test.hour)
		if len(rules) == 0 {
			t.Log("No rules returned for hour", test.hour)
			t.Fail()
		}
		first := rules[0]
		t.Log("Time:", test.hour, "EST. Expected first:", test.expected, "Actual first:", first.Value)

		if first.Value != test.expected {
			t.Log("failed", first.Value)
			t.Fail()
		}
	}
}

func TestRunStage(t *testing.T) {
	tests := []struct {
		hour     int
		position int
		find     string
		outcome  bool
	}{
		{hour: 10, position: 64, find: "'Plong'", outcome: true},
		{hour: 10, position: 66, find: "'Pling'", outcome: true},
		{hour: 12, position: 60, find: "'FizzBuzz'", outcome: true},
		{hour: 15, position: 57, find: "'Fizz'", outcome: true},
		{hour: 5, position: 5, find: "'5'", outcome: true},
		{hour: 15, position: 5, find: "'Buzz'", outcome: true},
		{hour: 15, position: 6, find: "'Fizz'", outcome: true},
		{hour: 15, position: 17, find: "'18'", outcome: false}, // we don't expect 18 to output in the 17th position
	}

	for _, test := range tests {
		stg := process.GetRules(test.hour)
		out := process.RunStage(1, 100, stg)
		res := out[test.position-1]

		t.Log("Time:", test.hour, "EST. Position", test.position, "looking for", test.find, "found", res, " expected outcome:", test.outcome)

		if strings.Contains(res, test.find) != test.outcome {
			t.Log("This was not an expected outcome")
			t.Fail()
		}
	}

}
