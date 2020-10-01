package main

import "fmt"

type animal interface {
	eat()
	move()
	speak()
}

type cow struct {
	name string
}

type bird struct {
	name string
}

type snake struct {
	name string
}

func (c cow) eat() {
	fmt.Printf("%s's food: grass\n", c.name)
}

func (c cow) move() {
	fmt.Printf("%s's locomotion: walk\n", c.name)
}

func (c cow) speak() {
	fmt.Printf("%s's spoken sound: moo\n", c.name)
}

func (b bird) eat() {
	fmt.Printf("%s's food: worms\n", b.name)
}

func (b bird) move() {
	fmt.Printf("%s's locomotion: fly\n", b.name)
}

func (b bird) speak() {
	fmt.Printf("%s's spoken sound: peep\n", b.name)
}

func (s snake) eat() {
	fmt.Printf("%s's food: mice\n", s.name)
}

func (s snake) move() {
	fmt.Printf("%s's locomotion: slither\n", s.name)
}

func (s snake) speak() {
	fmt.Printf("%s's spoken sound: hsss\n", s.name)
}

func main() {
	var cmd, s1, s2 string
	var c cow
	var b bird
	var s snake
	var created bool

	for {
		created = false
		fmt.Printf("Enter your command > ")
		_, _ = fmt.Scan(&cmd, &s1, &s2)

		if cmd == "newanimal" {
			switch s2 {
			case "cow":
				c = cow{s1}
				created = true
				// fmt.Println(c)
			case "bird":
				b = bird{s1}
				created = true
				// fmt.Println(b)
			case "snake":
				s = snake{s1}
				created = true
				// fmt.Println(s)
			default:
				fmt.Println("Undefined animal type!")
			}

			if created {
				fmt.Println("Created it!")
			}
		} else if cmd == "query" {
			switch s1 {
			case c.name:
				printQuery(c, s2)
			case b.name:
				printQuery(b, s2)
			case s.name:
				printQuery(s, s2)
			default:
				fmt.Println("Name not found!")
			}
		} else {
			fmt.Println("Wrong command!")
		}

	}
}

func printQuery(anim animal, action string) {
	switch action {
	case "eat":
		anim.eat()
	case "move":
		anim.move()
	case "speak":
		anim.speak()
	default:
		fmt.Println("Wrong input for action query!")
	}
}
