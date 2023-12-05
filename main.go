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
	p := buildProblem("Day 1", 1, 1, dayOnePartOne)
	runResults(p)
}

func buildProblem(name string, day int, num int, method func() int) problem {
	return problem{
		name:   name,
		day:    day,
		num:    num,
		method: method,
	}
}

func runResults(p problem) {
	start := time.Now()
	result := p.method()
	duration := time.Since(start)

	log.Printf("%v: result = %v, duration = %v", p.name, result, duration)
}
