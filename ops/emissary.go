package ops

import (
	"github.com/advanced-go/agency/common"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/resiliency/guidance"
)

type initOfficer func(origin core.Origin, handler messaging.OpsAgent) messaging.OpsAgent

// emissary attention
func emissaryAttend[T messaging.Notifier](agent *ops, initAgent initOfficer) {
	var notify T

	for {
		select {
		case msg := <-agent.emissary.C:
			switch msg.Event() {
			case messaging.ShutdownEvent:
				shutdown(agent)
				notify.OnMessage(agent, msg, agent.emissary)
				return
			case messaging.DataChangeEvent:
				if msg.IsContentType(guidance.ContentTypeCalendar) {
					agent.caseOfficers.Broadcast(msg)
				}
				notify.OnMessage(agent, msg, agent.emissary)
			case stopAgents:
				agent.caseOfficers.Shutdown()
				notify.OnMessage(agent, msg, agent.emissary)
			case startAgents:
				if agent.caseOfficers.Count() == 0 {
					initialize[T](agent, initAgent)
				}
				notify.OnMessage(agent, msg, agent.emissary)
			default:
				notify.OnError(agent, agent.Handle(common.MessageEventErrorStatus(agent.agentId, msg)))
			}
		default:
		}
	}
}

func initialize[T messaging.Notifier](o *ops, agent initOfficer) {
	var t T

	a := agent(westOrigin, o)
	err := o.caseOfficers.Register(a)
	if err != nil {
		t.OnError(o, o.Handle(core.NewStatusError(core.StatusInvalidArgument, err)))
	} else {
		a.Run()
	}
	a = agent(centralOrigin, o)
	err = o.caseOfficers.Register(a)
	if err != nil {
		t.OnError(o, o.Handle(core.NewStatusError(core.StatusInvalidArgument, err)))
	} else {
		a.Run()
	}
}
