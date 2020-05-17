package main

import (
	"fmt"
	"os"
	"flag"
	"encoding/csv"
	"strings"
	"time"
	"strconv"
)

type problem struct {
	q string
	a string
}

func main() {
	
	filename := flag.String("filename", "problems.csv", "pathname of csv file representing the quiz user will take")
	d := flag.Int("duration", 2, "integer representing the number of seconds to wait before the time runs out")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("attempted to open file with filename %v but could not\n", *filename)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	tuples, err := r.ReadAll()
	if err != nil {
		fmt.Printf("attempted to read file with filename %v but could not\n", *filename)
		os.Exit(1)
	}

	duration, err := time.ParseDuration((strconv.Itoa(*d) + "s"))
	if err != nil {
		fmt.Println(err)
		fmt.Printf("attempted to receive duration, %d, but was not an valid integer\n", *d)
		os.Exit(1)
	}

	var count int
	end := func() {
		fmt.Printf("\nScore: %d of %d\n", count, len(tuples))
		os.Exit(0)
	}

	quiz := make([]problem, len(tuples))
	for i, tuple := range tuples {
		quiz[i] = problem{
			q: strings.TrimSpace(tuple[0]),
			a: strings.TrimSpace(tuple[1]),
		}
	}

	for _, problem := range quiz {
		t := time.NewTimer(duration)
		fmt.Printf("%v: ", problem.q)
		ansChan := make(chan string)
		go func() {
			var attempt string
			fmt.Scanf("%s\n", &attempt)
			ansChan  <-attempt
		}()
		select {
			case <- t.C:
				end()
				return;
			case attempt := <-ansChan:
				if attempt == problem.a {
					count++
				} else {
					end()
				}
		}
	}
	fmt.Printf("Perfect!")
	end()
}

