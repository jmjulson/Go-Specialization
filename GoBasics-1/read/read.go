package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Names struct {
	fname string
	lname string
}

func main() {
	ff := getFileName()
	
	readFile(ff)
}

func readFile(fileN string) {
	var txtlines []Names
	
	fn := fileN + ".txt"
	file, err := os.Open(fn)
 
	check(err)
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		//returns a slice
		a := strings.Split(scanner.Text(), " ")

		// create a struct for each name
		var name = Names{ a[0], a[1] }

		//add name struct to the Names struct
		txtLines := append(txtlines, name)

		//not a huge fan of this loop
		for _, eachline := range txtLines {
			fmt.Println(eachline.lname + ", " + eachline.fname)
		}
	}

	file.Close()
}

func check(err error){
	if err != nil{
		fmt.Errorf("error: %w", err)
	}
}

func getFileName() string {
	var fileName string
	fmt.Println("Enter Name of file (exclude file extension):")

	fmt.Scanln(&fileName)

	return fileName
}