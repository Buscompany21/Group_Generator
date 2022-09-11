package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	var studNames []string = readFile()

	var shuffleNames []string = shuffleNames(studNames)

	var numStudentsPerGroup int = promptStudentsPerGroup()

	var remainderOfStudents int = len(shuffleNames) % numStudentsPerGroup

	// Initial values of incrementers
	groupCounter := 1
	x := 1
	//to be able to tell where we are in the groups
	i := 1

	// Group creation
	for g := 0; g < len(shuffleNames)/numStudentsPerGroup; g++ {
		println("Group", groupCounter)

		if groupCounter <= remainderOfStudents {
			x = 0
		} else {
			x = 1
		}

		for x <= (numStudentsPerGroup) {
			fmt.Println(shuffleNames[i-1])
			x++
			i++
		}
		groupCounter++
	}
}

func promptStudentsPerGroup() int {
	fmt.Println("How many individuals do you want in each group?")

	var numStudentsInput int

	fmt.Scanln(&numStudentsInput)

	return numStudentsInput
}

func readFile() []string {
	var studentNames []string

	// open file
	file, err := os.Open("students.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// do something with a line
		studentNames = append(studentNames, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return studentNames
}

func shuffleNames(names []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})

	return names
}
