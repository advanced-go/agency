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

func (d *dispatchT) dispatch(agent *ops, event string) {
	switch event {
	case stopAgentsEvent:
		finalized := agent.caseOfficers.IsFinalized()
		fmt.Printf("test: dispatch(%v) -> [finalized:%v] [count:%v]\n", event, finalized, agent.caseOfficers.Count())
	case startAgentsEvent:
		finalized := agent.caseOfficers.IsFinalized()
		fmt.Printf("test: dispatch(%v) -> [finalized:%v] [count>0:%v\n", event, finalized, agent.caseOfficers.Count() > 0)
	case messaging.DataChangeEvent:
		opsAgent.Trace(agent, event, "Broadcast() -> calendar data change event")
	}
}
