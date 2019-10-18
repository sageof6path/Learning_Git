package external

import (
	"fmt"
)

type ExternalStruct struct {
	Y int
}

func (ex ExternalStruct) Try() {
	fmt.Print("external is working")
}
