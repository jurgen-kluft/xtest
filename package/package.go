package xtest

import (
	"github.com/jurgen-kluft/xbase/package"
	"github.com/jurgen-kluft/xcode/denv"
)

// GetPackage returns the package object of 'xtest'
func GetPackage() *denv.Package {
	// dependencies
	xbasepkg := xbase.GetPackage()

	// 'xtest' application
	mainapp := denv.SetupDefaultCppAppProject("xtest", "github.com\\jurgen-kluft\\xtest")
	mainapp.Dependencies = append(mainapp.Dependencies, xbasepkg.GetMainLib())

	mainpkg := denv.NewPackage("xtest")
	mainpkg.AddMainApp(mainapp)
	return mainpkg
}
