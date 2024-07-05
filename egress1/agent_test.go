package egress1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
)

func ExampleAgentUri() {
	origin := core.Origin{
		Region:     "us-central1",
		Zone:       "c",
		SubZone:    "sub-zone",
		Host:       "host",
		InstanceId: "",
	}
	u := AgentUri(origin)
	fmt.Printf("test: AgentUri() -> [%v]\n", u)

	origin.Region = "us-west1"
	origin.Zone = "a"
	origin.SubZone = ""
	u = AgentUri(origin)
	fmt.Printf("test: AgentUri() -> [%v]\n", u)

	//Output:
	//test: AgentUri() -> [egress-controller1:us-central1.c.sub-zone.host]
	//test: AgentUri() -> [egress-controller1:us-west1.a.host]

}
