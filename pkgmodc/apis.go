package pkgmodc

import (
	"fmt"

	"github.com/EnricoPicci/go-class-dep-mgmt-mod-c/intapis"

	"github.com/EnricoPicci/go-class-dep-mgmt-mod-d/pkgmodd"
)

func ApiThatCallsInternalApi() {
	fmt.Println("I am an API of module C that calls an internal API of module C")
	intapis.InternalAPI()
}

func ApiThatCallsModuleDApi() {
	fmt.Println("I am an API of module C that calls an an API of module D")
	pkgmodd.DoStuffModuleD()
}
