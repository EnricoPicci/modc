package pkgmodc1

import (
	"fmt"

	"github.com/EnricoPicci/modc/intapis"
)

func ApiThatCallsInternalApi() {
	fmt.Println("I am an API of module C that calls an internal API of module C")
	intapis.InternalAPI()
}
