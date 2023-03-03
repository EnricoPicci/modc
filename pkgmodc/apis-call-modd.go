package pkgmodc

import (
	"fmt"

	"github.com/EnricoPicci/modd/pkgmodd"
)

func ApiThatCallsModuleDApi() {
	fmt.Println("I am an API of module C that calls an an API of module D")
	pkgmodd.DoStuffModuleD()
}
