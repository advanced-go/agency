package ops

import "github.com/advanced-go/common/messaging"

type dispatcher interface {
	dispatch(ops messaging.OpsAgent, agent messaging.Agent, event string)
}

type dispatch struct{}

func newDispatcher() dispatcher {
	d := new(dispatch)
	return d
}

func (d *dispatch) dispatch(ops messaging.OpsAgent, agent messaging.Agent, event string) {
	switch event {
	case stopAgentsEvent:
		opsAgent.Trace(agent, event, "stopping case officer agents")
	case startAgentsEvent:
		opsAgent.Trace(agent, event, "starting case officer agents")
	}
}
