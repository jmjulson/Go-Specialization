package main 

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	var s = getInput()
	result(s)
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	stringIn, _ := reader.ReadString('\n')
	
	return stringIn
}

func result(enteredString string){
	var cleanUp string = strings.ToLower(enteredString)
	var result string = strings.TrimSpace(cleanUp)
	fmt.Println("Some Touch-Up:", result)
	
	if strings.HasPrefix(result, "i") && strings.Contains(result, "a") && strings.HasSuffix(result, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
