package Tests

import (
	"fmt"

	. "github.com/Lucky31415/GoProjektarbeit/Expressions"
	. "github.com/Lucky31415/GoProjektarbeit/Optional"
	. "github.com/Lucky31415/GoProjektarbeit/Parser"
)

func display(e Optional[Expression]) {
	if e.IsNothing() {
		fmt.Println("nothing")
	} else {
		fmt.Println(e.FromJust().Pretty())
	}
}

func TestParserGood() {
	display(NewParser("(1)").Parse())
	display(NewParser("1").Parse())
	display(NewParser("1 + 0").Parse())
	display(NewParser("1 + (0) ").Parse())
	display(NewParser("1 + 2 * 0 ").Parse())
	display(NewParser("1 * 2 + 0 ").Parse())

	display(NewParser("(1 + 2) * 0 ").Parse())
	display(NewParser("(1 + 2) * 0 + 2").Parse())
}
