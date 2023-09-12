package main

import "fmt"

type pilot struct {
	Name     string
	Life     float32
	Age      int
	Aircraft int
}

func main() {
	const AIRCRAFT1 = 1
	donnie := pilot{}
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}
