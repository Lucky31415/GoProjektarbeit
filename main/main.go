package main

import "fmt"

type RZ struct {
	Zaehler int
	Nenner  int
}

func printRZ(x RZ) {
	fmt.Printf("%+v/%+v", x.Zaehler, x.Nenner)
}

func main() {
	x := RZ{
		Zaehler: 5,
		Nenner:  7,
	}

	printRZ(x)
}
