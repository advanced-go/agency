package ops

import (
	"fmt"
	"github.com/advanced-go/common/messaging"
)

type dispatchT struct{}

func newTestDispatcher() dispatcher {
	d := new(dispatchT)
	return d
}

func (d *dispatchT) setup(_ *ops, _ string) {}

func (d *dispatchT) dispatch(agent *ops, event string) {
	switch event {
	case stopAgentsEvent:
		fmt.Printf("test: dispatch(%v) -> [count:%v]\n", event, agent.caseOfficers.Count())
	case startAgentsEvent:
		fmt.Printf("test: dispatch(%v) -> [count>0:%v]\n", event, agent.caseOfficers.Count() > 0)
	case messaging.DataChangeEvent:
		agent.Trace(agent, messaging.EmissaryChannel, event, "Broadcast() -> calendar data change event")
	}
}

func ExampleTestDispatcher() {
	fmt.Printf("test: TestDispatch() \n")

	//Output:
	//test: TestDispatch()

}
