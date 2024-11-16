package ops

import (
	"github.com/advanced-go/agency/common"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/resiliency/guidance"
)

type initOfficer func(origin core.Origin, handler messaging.OpsAgent) messaging.OpsAgent

// emissary attention
func emissaryAttend(agent *ops, initAgent initOfficer) {
	for {
		select {
		case msg := <-agent.emissary.C:
			switch msg.Event() {
			case messaging.ShutdownEvent:
				agent.shutdown()
				return
			case messaging.DataChangeEvent:
				if msg.IsContentType(guidance.ContentTypeCalendar) {
					agent.caseOfficers.Broadcast(msg)
				}
			case stopAgents:
				agent.caseOfficers.Shutdown()
			case startAgents:
				if agent.caseOfficers.Count() == 0 {
					initialize(agent, initAgent)
				}
			default:
				agent.Handle(common.MessageEventErrorStatus(agent.agentId, msg))
			}
		default:
		}
	}
}

func initialize(o *ops, agent initOfficer) {
	a := agent(westOrigin, o)
	err := o.caseOfficers.Register(a)
	if err != nil {
		o.Handle(core.NewStatusError(core.StatusInvalidArgument, err))
	} else {
		a.Run()
	}
	a = agent(centralOrigin, o)
	err = o.caseOfficers.Register(a)
	if err != nil {
		o.Handle(core.NewStatusError(core.StatusInvalidArgument, err))
	} else {
		a.Run()
	}
}
