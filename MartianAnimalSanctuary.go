package main

import (
	"fmt"
	"math/rand"
	"time"
)

var names = []string{"River", "Tiffany", "Kayina", "Sarah", "Soraka", "Star", "DTM", "Gegi", "Pancake", "Sgt"}

type animal interface {
	String() string
	move() string
	eat() string
}

type martian struct {
	name     string
	foods    []string
	movement string
}

type animeFemboy struct {
	martian
}

type vampireFemboy struct {
	martian
}

type thickMushroom struct {
	martian
}

func newAnimeFemboy() animeFemboy {
	toReturn := animeFemboy{}
	toReturn.name = names[rand.Intn(len(names))]
	toReturn.foods = []string{"carrot", "banana", "eggplant", "zucchini", "cucumber", "yogurt", "bubble tea"}
	toReturn.movement = "skips around happily and swings their arms aesthetically in an attempt to attract mates"
	return toReturn
}

func newVampireFemboy() vampireFemboy {
	toReturn := vampireFemboy{}
	toReturn.name = names[rand.Intn(len(names))]
	toReturn.foods = []string{"neck", "feet", "collar", "thigh"}
	toReturn.movement = "sleuths around stealthily looking for their next prey"
	return toReturn
}

func newThickMushroom() thickMushroom {
	toReturn := thickMushroom{}
	toReturn.name = names[rand.Intn(len(names))]
	toReturn.foods = []string{"water", "sweat"}
	toReturn.movement = "grows taller and thicker to attract mates"
	return toReturn
}

func (f martian) move() string {
	return f.movement
}

func (f martian) eat() string {
	return f.foods[rand.Intn(len(f.foods))]
}

func (f martian) String() string {
	return f.name
}

func runSancturary() {
	rand.Seed(time.Now().Unix())

	sancturary := make([]animal, 0, 10)
	for i := 0; i < 10; i++ {
		decider := rand.Intn(3)
		switch decider {
		case 0:
			sancturary = append(sancturary, newAnimeFemboy())
		case 1:
			sancturary = append(sancturary, newVampireFemboy())
		case 2:
			sancturary = append(sancturary, newThickMushroom())
		}
	}

	fmt.Printf("Starting Sanctruary Simulation...\n")
	for day := 0; day < 2; day++ {
		fmt.Println("________________________________")
		fmt.Printf("DAY %d\n\n", day)
		for hour := 0; hour < 12; hour++ {
			fmt.Printf("Day, Hour: %d\n", hour)
			animalIndex := rand.Intn(len(sancturary))
			animalAction := rand.Intn(2)
			switch animalAction {
			case 0:
				fmt.Println(sancturary[animalIndex], sancturary[animalIndex].move())
			case 1:
				fmt.Println(sancturary[animalIndex], "chomps down on", sancturary[animalIndex].eat())
			}
			time.Sleep(time.Millisecond * 500)
		}
		for hour := 12; hour < 24; hour++ {
			fmt.Printf("Night, Hour: %d\n", hour)
			fmt.Println("Everyone is asleep")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
