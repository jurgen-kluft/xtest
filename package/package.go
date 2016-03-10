package xtest

import (
	"github.com/jurgen-kluft/xcode/denv"
)

// GetPackage returns the package object of 'xtest'
func GetPackage() *denv.Package {

	// 'xtest' library
	mainapp := denv.SetupDefaultCppAppProject("xtest", "github.com\\jurgen-kluft\\xtest")

	mainpkg := denv.NewPackage("xtest")
	mainpkg.AddMainApp(mainapp)
	return mainpkg
}
