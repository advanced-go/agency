package ops

import "github.com/advanced-go/common/messaging"

type dispatcher interface {
	setup(agent *ops, event string)
	dispatch(agent *ops, event string)
}

type dispatch struct{}

func newDispatcher() dispatcher {
	d := new(dispatch)
	return d
}

func (d *dispatch) setup(_ *ops, _ string) {}

func (d *dispatch) dispatch(agent *ops, event string) {
	switch event {
	case messaging.StartupEvent:
		agent.Trace(agent, messaging.EmissaryChannel, event, "startup")
	case messaging.ShutdownEvent:
		agent.Trace(agent, messaging.EmissaryChannel, event, "shutdown")
	case stopAgentsEvent:
		agent.Trace(agent, messaging.EmissaryChannel, event, "stopping case officer agents")
	case startAgentsEvent:
		agent.Trace(agent, messaging.EmissaryChannel, event, "starting case officer agents")
	}
}
