package main

import "fmt"

func test1() {
	fmt.Printf("Test 1: \n")

	fmt.Printf("Eingabe 1: \n")
	rz1 := ScanRZ()

	fmt.Printf("Eingabe 2: \n")
	rz2 := ScanRZ()

	PrintRZ(rz1)
	PrintRZ(rz2)

	fmt.Printf("Add Ausgabe: \n")
	PrintRZ(AddRz(rz1, rz2))

	fmt.Printf("Mult Ausgabe: \n")
	PrintRZ(MultRZ(rz1, rz2))
}
