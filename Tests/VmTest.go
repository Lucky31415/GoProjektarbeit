package Tests

import (
	"fmt"

	. "github.com/Lucky31415/GoProjektarbeit/Optional"
	. "github.com/Lucky31415/GoProjektarbeit/VM"
)

func showVMRes(r Optional[int]) {
	if r.IsNothing() {
		fmt.Println("VM stack (top): empty")
	}
	fmt.Println("VM stack (top): " + fmt.Sprint(r.FromJust()))
}

func TestVM() {
	codes := []Code{
		NewPush(1),
		NewPush(2),
		NewPush(3),
		NewMult(),
		NewPlus(),
	}
	vm := NewVm(codes)
	res := vm.Run()
	showVMRes(res)

	codes = []Code{
		NewPush(2),
		NewPush(3),
		NewPush(5),
		NewPlus(),
		NewMult(),
	}
	vm = NewVm(codes)
	res = vm.Run()
	showVMRes(res)
}
