package turingmachine

import (
	"fmt" 		// for Println
	"bufio"		// for Scanner
	"os"		// for Open
	"log"		// for log
	//"strings"	// for strings
)

type Tuple struct {
	output, next, direction string
}

type TuringMachine struct {
	rules map[string]map[string]Tuple 	// dict[string]->dict[string]->tuple
	state string 						// current state
	head uint64 						// for reading the tape
}

func NewTuringMachine(filename string) *TuringMachine {
	t := new(TuringMachine) // new returns a ptr to a TuringMachine type
	t.head = 1 // set our head to 1 to skip the first * on the tape

	// open file for read access
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create a new scanner
	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanWords)

	// scan the file line by line and store it into a slice
	for scanner.Scan() {
		//line := strings.Split(scanner.Text(), "\n")

		// get the start state
		if t.state == "" {
			t.state = string(scanner.Text()[0])
		}

		// extract each rule
		state = string(scanner.Text()[0])
		input = string(scanner.Text()[1])
		output = string(scanner.Text()[2])
		next = string(scanner.Text()[3])
		direction = string(scanner.Text()[4])
		t.rules[scanner.Text()[0]][scanner.Text()[1]] = Tuple{scanner.Text()}
		//fmt.Println(line)
	}

	// check for scan errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	// create a slice of bytes to store file contents
	//data := make([]byte, 100)
	// read file contents into the slice
	//t.state = start
	fmt.Println(t.state) // for testing
	return t
}
