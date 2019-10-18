package main

import (
	"gitDownloads/LearningGit/external"
	"gitDownloads/LearningGit/internal"
)

func main() {
	ext := external.ExternalStruct{Y: 10}
	inter := internal.InternalStruct{X: 20}
	ext.Try(inter)
}
