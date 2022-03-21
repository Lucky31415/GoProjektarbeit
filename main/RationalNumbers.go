package main

import "fmt"

type RZ struct {
	Zaehler int
	Nenner  int
}

func PrintRZ(x RZ) {
	fmt.Printf("%+v/%+v", x.Zaehler, x.Nenner)
}
