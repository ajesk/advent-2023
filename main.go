package main

import (
	"log"
	"time"
)

type problem struct {
	name   string
	day    int
	num    int
	method func() int
}

func main() {
	problems := []problem{}

	problems = append(problems, buildProblem("Trebuchet?!", 1, 1, dayOnePartOne))
	problems = append(problems, buildProblem("Calibration", 1, 2, dayOnePartTwo))
	problems = append(problems, buildProblem("Cube Conundrum", 2, 1, dayTwoPartOne))
	problems = append(problems, buildProblem("Fewest Cubes", 2, 2, dayTwoPartTwo))
	problems = append(problems, buildProblem("Gear Ratios", 3, 1, dayThreePartOne))
	problems = append(problems, buildProblem("The Actual Gears", 3, 2, dayThreePartTwo))
	problems = append(problems, buildProblem("Scratchcards", 4, 1, dayFourPartOne))
	// problems = append(problems, buildProblem("Exponential Scratchcards", 4, 2, dayFourPartTwo))
	problems = append(problems, buildProblem("Min of Seeds", 5, 1, day5Part1))
	problems = append(problems, buildProblem("Min of Seed Ranges", 5, 2, day5Part2))
	runResults(problems)
}

func buildProblem(name string, day int, num int, method func() int) problem {
	return problem{
		name:   name,
		day:    day,
		num:    num,
		method: method,
	}
}

func runResults(problems []problem) {

	for _, prob := range problems {
		start := time.Now()
		result := "redacted"
		prob.method()
		duration := time.Since(start)

		log.Printf(
			"[Day %v, Part %v, %v]: result = %v, duration = %v",
			prob.day, prob.num, prob.name, result, duration,
		)
	}
}
