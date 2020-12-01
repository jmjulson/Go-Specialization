package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

type Incoming struct {
	creature string
	action   string
}

func main() {
	cow := Animal{"grass", "walk", "mooo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for i := 0; i >= 0; i++ {
		entered := getAnimalInput()

		if entered.creature == "cow" {
			doSomething(cow, entered.action)
		} else if entered.creature == "bird" {
			doSomething(bird, entered.action)
		} else if entered.creature == "snake" {
			doSomething(snake, entered.action)
		} else {
			fmt.Println("*** Wrong Animal ***")
			break
		}
	}
}

func doSomething(a Animal, action string) {
	switch {
	case action == "eat":
		a.Eat()
	case action == "move":
		a.Move()
	case action == "speak":
		a.Speak()
	default:
		fmt.Println("*** Wrong Action ***")
		break
	}
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func getAnimalInput() Incoming {
	var input string
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an Animal (bird OR cow OR snake) & requested action (eat OR speak OR move)")
	fmt.Println(" \nYour response should be sperated by a space i.e \"bird move\"--> ")

	input, _ = scanner.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	str := strings.Split(input, " ")
	var in = Incoming{str[0], str[1]}

	return in
}
