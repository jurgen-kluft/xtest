package main

import (
	"github.com/jurgen-kluft/xtest/package"
	"github.com/jurgen-kluft/xcode"
)

func main() {
	xcode.Generate(xtest.GetProject())
}
