package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	name := getName()
	address := getAddress()

	pMap := map[string] string {
		"name" : name,
		"address" : address,
	} 

	foo, err := json.Marshal(pMap)

	if err != nil {
		fmt.Errorf("error: %w", err)
	}

	fmt.Println(foo)
}

func getAddress() string{
	var address string

	fmt.Println("Enter an Addres:")
	fmt.Scanln(&address)

	return address
}

func getName() string{
	var name string

	fmt.Println("Enter a Name:")
	fmt.Scanln(&name)

	return name
}
