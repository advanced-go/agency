package ops

import "github.com/advanced-go/common/messaging"

type dispatchT struct{}

func newTestDispatcher() dispatcher {
	d := new(dispatchT)
	return d
}

func (d *dispatchT) dispatch(ops messaging.OpsAgent, agent messaging.Agent, event string) {
	switch event {
	case stopAgentsEvent:
		//opsAgent.Trace(agent, event, "stopping case officer agents")
	case startAgentsEvent:
		//opsAgent.Trace(agent, event, "starting case officer agents")
	case messaging.DataChangeEvent:
		opsAgent.Trace(agent, event, "data change Broadcast()")
	}
}
