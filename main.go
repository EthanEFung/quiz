package main

import (
	"fmt"
	"io"
	"log"
	"io/ioutil"
	"encoding/csv"
	"strings"
	"bufio"
	"os"
	"flag"
)

func main() {
	filename := flag.String("filename", "problems.csv", "csv file of the quiz to be taken")
	flag.Parse()
	fmt.Println(*filename)
	content, err := ioutil.ReadFile(*filename)
	if  err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(content)))
	records, err := r.ReadAll()
	var count int
	for _, record := range records {
		question, ans := record[0], record[1]
		ans = strings.TrimSpace(ans)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(question + " =")
		buf := bufio.NewReader(os.Stdin)
		sentence, e := buf.ReadBytes('\n')
		if e != nil {
			log.Fatal(e)
		} 
		input := string(sentence)
		input = strings.TrimSpace(input)
		
		if input == ans {
			count++
		} else {
			break
		}
	}
	fmt.Printf("%v of %v correct", count, len(records))
}
