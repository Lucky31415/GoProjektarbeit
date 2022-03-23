package main

import (
	"fmt"
	//"github.com/Lucky31415/GoProjektarbeit/Expressions"
)

func main() {

	me := NewMultExp(NewIntExp(1), NewPlusExp(NewIntExp(3), NewIntExp(5)))
	fmt.Printf("%s = %d", me.Pretty(), me.Eval())

	//unitTest1()
	//test1()
	//rz1 := ScanRZ()
	//rz2 := ScanRZ()
	//
	//rzSum := AddRz(rz1, rz2)
	//PrintRZ(rzSum)

}
