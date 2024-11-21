package ops

import (
	"fmt"
	"github.com/advanced-go/common/host"
)

func ExampleStartupPing() {
	status := host.Ping(PkgPath)
	fmt.Printf("test: Ping() -> [status:%v]\n", status)

	//Output:
	//agency-ops : event:startup startuptest: Ping() -> [status:OK]

}
