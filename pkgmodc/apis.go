package pkgmodc

import (
	"fmt"

	"github.com/EnricoPicci/modc/intapis"
	"github.com/EnricoPicci/modd/pkgmodd"
)

func ApiThatCallsInternalApi() {
	fmt.Println("I am an API of module C that calls an internal API of module C")
	intapis.InternalAPI()
}

func ApiThatCallsModuleDApi() {
	fmt.Println("I am an API of module C that calls an an API of module D")
	pkgmodd.DoStuffModuleD()
}
