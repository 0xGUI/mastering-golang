package main

import (

	"flag"
	"os"
	"fmt"
	"encoding/csv"
)

func main() {

	//present user problems
	//checks for input
	//checks if answer is correct
	//add point to user

	//get the file stdin arguments
	csvFilename := flag.String("csv", "questions.csv", "CSV File with questions")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	_ = file
	if err != nil {
	  exit(fmt.Sprintf ("Failed to open file: %s", *csvFilename))
	}

	//csv reader
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse csv file")
	}

	problems := parseLines(lines)

	for i, p := range problems {
	  fmt.Printf("Problem number %d : %s = \n", i+1, p.q)
	  var answer string
	  fmt.Scanf("%s\n", &answer)
	  if answer == p.a {
	    fmt.Printf("Correct!") 
    }
	}
}

// takes 2d lines string slice and return slice of problems
func parseLines (lines [][]string) []problem {
 ret := make ([]problem,len(lines))
 for i, line := range lines {
	ret[i] = problem{
	  q: line[0],
	  a: line[1],
	}
 }

 return ret
}

type problem struct {
	q string
	a string
}


func exit (err_msg string){
	fmt.Println (err_msg)
	os.Exit(1)
}
