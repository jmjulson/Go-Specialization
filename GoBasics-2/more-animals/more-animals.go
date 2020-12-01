package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	eat()      string
	speak() string
	move()    string
}

type Incoming struct {
	selection, name, thing string
}

type Cow struct {
	eat_, move_, speak_ string
}

type Bird struct {
	eat_, move_, speak_ string
}

type Snake struct {
	eat_, move_, speak_ string
}

func main() {
	fmt.Println("****************************************************")
	fmt.Println("* Please Enter one of the following:               *")
	fmt.Println("* (1) newanimal name animal [cow, snake, bird]     *")
	fmt.Println("* e.g. -> newanimal bob cow                        *")
	fmt.Println("* (2) -> query name action [speak, eat, move]      *")
	fmt.Println("* e.g. -> query bob speak                          *")
	fmt.Println("****************************************************")

	c := Cow{"grass", "walk", "mooo"}
	b := Bird{"worms", "fly", "peep"}
	s := Snake{"mice", "slither", "hsss"}
	listOfAnimals := make(map[string]Incoming)

	for i := 0; i >= 0; i++ {
		entry := getAnimalInput()

		switch {
		case entry.selection == "newanimal":
			listOfAnimals[entry.name] = Incoming{entry.selection, entry.name, entry.thing}
			fmt.Println("Created it!")
			fmt.Println("Running list of animals:")

			for _, value := range listOfAnimals {
				fmt.Println(value.name, "the", value.thing)
			}

		case entry.selection == "query":
			switch{
			case listOfAnimals[entry.name].thing == "cow":
				 animalAction(c, entry.thing, entry.name)
			case listOfAnimals[entry.name].thing == "bird":
				animalAction(b, entry.thing, entry.name)
			case listOfAnimals[entry.name].thing == "snake":
				animalAction(s, entry.thing, entry.name)
			default:
				fmt.Println("*** Please enter a valid animal (bird, cow, snake) ***")
				break
			}
		default:
			fmt.Println("*** Please enter newanimal or query ***")
			break
		}
	}
}

func (c Cow) eat() string{
	return c.eat_
}

func (c Cow) speak() string{
	return c.speak_
} 

func (c Cow) move() string{
	return c.move_
} 

func (c Bird) eat() string{
	return c.eat_
}

func (c Bird) speak() string{
	return c.speak_
} 

func (c Bird) move() string{
	return c.move_
}

func (c Snake) eat() string{
	return c.eat_
}

func (c Snake) speak() string{
	return c.speak_
} 

func (c Snake) move() string{
	return c.move_
}

func getAnimalInput() Incoming {
	var input string
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("-> ")

	input, _ = scanner.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	str := strings.Split(input, " ")

	if (len(str) < 3){
		fmt.Print("*** Missing Parameters ***")
	}

	var in = Incoming{str[0], str[1], str[2]}

	return in
}

func animalAction(a Animal, thing string, name string) {
    switch {
	case thing == "eat":
			fmt.Println(name, "likes to eat", a.eat())
	case thing == "speak":
		fmt.Println(name, "makes a", a.speak(), "sound")
	case thing == "move":
		fmt.Println(name, "gets around with a", a.move())
	default:
		fmt.Println("*** Ooops!", thing, "is not a valid action ***")
	}

}
