package external

import (
	"fmt"
	"https://github.com/sageof6path/Learning_Git/internal"
)

type ExternalStruct struct {
	Y int
}

func (ex ExternalStruct) Try(internalStruct internal.InternalStruct)  {
	fmt.Print("external is working")
}

