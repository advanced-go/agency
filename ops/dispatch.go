package ops

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
	case stopAgentsEvent:
		agent.Trace(agent, event, "stopping case officer agents")
	case startAgentsEvent:
		agent.Trace(agent, event, "starting case officer agents")
	}
}
