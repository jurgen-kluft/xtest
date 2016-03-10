package main

import (
	"github.com/jurgen-kluft/xcode"
	"github.com/jurgen-kluft/xtest/package"
)

func main() {
	xcode.Generate(xtest.GetPackage())
}
