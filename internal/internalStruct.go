package internal

import (
	"fmt"
	"gitDownloads/LearningGit/external"
)

type InternalStruct struct {
	X int
}

func (x InternalStruct) Try(externalStruct external.ExternalStruct) {
	fmt.Print("its working")
}
