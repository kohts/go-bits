package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var people []Person

	var fname string
	fmt.Printf("Enter filename: ")
	fmt.Scan(&fname)

	file, err := os.Open(fname)
	if err != nil {
		log.Fatal("Can't open the file %v for reading: %v", fname, err)
	}
	defer file.Close()

	line := ""
	reader := bufio.NewReader(file)

	for {
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fields := strings.Fields(string(line[:]))
		if len(fields) > 0 {
			people = append(people, Person{fname: "", lname: ""})

			people[len(people)-1].fname = fields[0]
			if len(fields[0]) > 20 {
				people[len(people)-1].fname = fields[0][:20]
			}

			if len(fields) > 1 {
				people[len(people)-1].lname = fields[1]
				if len(fields[1]) > 20 {
					people[len(people)-1].lname = fields[1][:20]
				}
			}
		}
	}
	file.Close()

	for i, p := range people {
		fmt.Printf("person %v, first name: %v, last name: %v\n", i+1, p.fname, p.lname)
	}

}
