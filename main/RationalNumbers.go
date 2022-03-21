package main

import "fmt"

type RZ struct {
	Zaehler int
	Nenner  int
}

func PrintRZ(x RZ) {
	fmt.Printf("%+v/%+v\n", x.Zaehler, x.Nenner)
}

func ScanRZ() RZ {
	var zaehler int
	var nenner int

	fmt.Scanf("%d/%d\n", &zaehler, &nenner)

	x := RZ{
		Zaehler: zaehler,
		Nenner:  nenner,
	}

	return x
}

func AddRz(x RZ, y RZ) RZ {
	var xTemp, yTemp RZ

	xTemp = SkaliereRZ(x, y.Nenner)
	yTemp = SkaliereRZ(y, x.Nenner)

	z := RZ{
		Zaehler: xTemp.Zaehler + yTemp.Zaehler,
		Nenner:  xTemp.Nenner,
	}

	return z
}

func AddRz2(x RZ, y RZ) RZ {
	SkaliereRZ2(&x, y.Nenner)
	SkaliereRZ2(&y, x.Nenner)

	z := RZ{
		Zaehler: x.Zaehler + y.Zaehler,
		Nenner:  x.Nenner,
	}

	return z
}

func SkaliereRZ(x RZ, k int) RZ {
	z := RZ{
		Zaehler: x.Zaehler * k,
		Nenner:  x.Nenner * k,
	}

	return z
}

func SkaliereRZ2(x *RZ, k int) {
	x.Zaehler *= k
	x.Nenner *= k
}

func MultRZ(x RZ, y RZ) RZ {
	z := RZ{
		Zaehler: x.Zaehler * y.Zaehler,
		Nenner:  x.Nenner * y.Nenner,
	}

	return z
}
