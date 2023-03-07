package pkgmodc2

import (
	"fmt"

	"github.com/EnricoPicci/modd/pkgmodd"
)

func ApiThatCallsModuleDApi() {
	fmt.Println("I am an API of module C that calls an an API of an external module, module D")
	pkgmodd.DoStuffModuleD()
}
