package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (anim *Animal) InitAnimal(f, l, n string) {
	anim.food = f
	anim.locomotion = l
	anim.noise = n
}

func (anim *Animal) Eat() {
	fmt.Println("food: ", anim.food)
}

func (anim *Animal) Move() {
	fmt.Println("locomotion: ", anim.locomotion)
}

func (anim *Animal) Speak() {
	fmt.Println("spoken sound:", anim.noise)
}

func main() {
	cow := new(Animal)
	bird := new(Animal)
	snake := new(Animal)

	cow.InitAnimal("grass", "walk", "moo")
	bird.InitAnimal("worms", "fly", "peep")
	snake.InitAnimal("mice", "slither", "hss")

	var reqAnim, reqAction string
	for {
		fmt.Printf("Enter your request > ")
		_, _ = fmt.Scan(&reqAnim, &reqAction)

		switch reqAnim {
		case "cow":
			fmt.Printf("Cow's ")
			printAction(cow, reqAction)
		case "bird":
			fmt.Printf("Bird's ")
			printAction(bird, reqAction)
		case "snake":
			fmt.Printf("Snake's ")
			printAction(snake, reqAction)
		default:
			fmt.Println("Invalid animal name! Check your request please.")
		}
	}
}

func printAction(anim *Animal, action string) {
	switch action {
	case "eat":
		anim.Eat()
	case "move":
		anim.Move()
	case "speak":
		anim.Speak()
	default:
		fmt.Println("Invalid action name! Check you request please.")
	}
}
