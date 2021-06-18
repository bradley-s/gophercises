package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Parser interface {
	Parse(s string) map[string]string
}

type DefaultParser struct{}

func NewDefaultParser() DefaultParser {
	return DefaultParser{}
}

func (p DefaultParser) Parse(f string) map[string]string {
	parsedData := make(map[string]string)

	file, err := os.Open(f)
	if err != nil {
		fmt.Sprintf("Failed to open the CSV file: %s\n", f)
		os.Exit(1)
	}

	r := csv.NewReader(file)

	data, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data {
		parsedData[v[0]] = v[1]
	}
	return parsedData
}
