package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	cow := animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	bird := animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	snake := animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	mapping := map[string]animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	i := 0
	animals := make([]string, len(mapping))
	for k := range mapping {
		animals[i] = k
		i++
	}
	sort.Strings(animals)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter animal name and requested info about it (type 'exit' to quit).")
	fmt.Printf("Animals available: %v\n", animals)

	for {
		fmt.Printf("> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		slice := strings.Split(input, " ")
		if len(slice) != 2 {
			fmt.Println("Please enter only two words, animal name and action.")
			continue
		}

		animalName := slice[0]
		index := sort.SearchStrings(animals, animalName)
		if animals[index] != animalName {
			fmt.Println("Unknown animal.")
			continue
		}

		animal := mapping[slice[0]]
		action := strings.ToLower(slice[1])

		if action == "eat" {
			animal.eat()
		} else if action == "move" {
			animal.move()
		} else if action == "speak" {
			animal.speak()
		} else {
			fmt.Println("Undefined action, please choose from 'eat', 'move' or 'speak'.")
		}
	}
}

func (a *animal) eat() {
	fmt.Println(a.food)
}

func (a *animal) move() {
	fmt.Println(a.locomotion)
}

func (a *animal) speak() {
	fmt.Println(a.noise)
}
