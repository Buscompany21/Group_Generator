package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {


    // prompt user for number of students in each group

    fmt.Println("How many students should be in each group? ")
    var studNum int 

    // assign input to variable
    fmt.Scanln(&studNum)

    // test to see if input is int
    fmt.Print(studNum * 5)


    // open file
    f, err := os.Open("IS405-562 - Section002.txt")
    if err != nil {
        log.Fatal(err)
    }
    // close the file
    defer f.Close()

    // read the file
    scanner := bufio.NewScanner(f)

	names := []string{}


    for scanner.Scan() {
        // append each line to array
        names = append(names, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // print array
	fmt.Printf("%v", names)

}
