package main

import "fmt"

type Hobby int

//go:generate go run golang.org/x/tools/cmd/stringer -type Hobby
const (
	Undefined Hobby = iota
	Diving
	Motorcycles
	Espresso
	MountainBiking
	Skiing
)

func main() {
	fmt.Printf("One of my hobbies is: %s\n", Diving)
}
