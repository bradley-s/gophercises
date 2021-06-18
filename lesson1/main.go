package main

import (
	"flag"
	"gophercises.com/v2/lesson1/csv"
	"gophercises.com/v2/lesson1/quiz"
)

var (
	filename string
	limit    int
)

const DefaultLimit = 30

func main() {
	problems := processInput()

	handleQuiz(problems)
}

func handleQuiz(problems map[string]string) {
	handler := quiz.NewDefaultHandler(limit)

	handler.Handle(problems)
}

func processInput() map[string]string {
	parser := csv.NewDefaultParser()

	flag.StringVar(&filename, "csv", "/Users/brad/go/src/github.com/gophercises/lesson1/problems.csv", `a csv file in the format of 'question, answer'`)
	flag.IntVar(&limit, "limit", DefaultLimit, `the time limit for the quiz in secs`)
	flag.Parse()

	return parser.Parse(filename)
}
