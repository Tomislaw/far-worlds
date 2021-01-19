package main

import (
	"fmt"

	"tostanie.com/globalworld/ecs"
)

type Walk struct {
	Direction string
	Distance  float64
}

// The Talk component
type Talk struct {
	Message string
}

func main() {
	fmt.Println("Hello, 世界")

	manager := ecs.Manager{}

}
