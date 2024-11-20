package ops

import (
	"errors"
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
			agent.OnMessage(agent, msg, agent.emissary)
			switch msg.Event() {
			case messaging.ShutdownEvent:
				agent.finalize()
				return
			case messaging.DataChangeEvent:
				if msg.IsContentType(guidance.ContentTypeCalendar) {
					agent.caseOfficers.Broadcast(msg)
					agent.OnTrace(agent, "officers.Broadcast()")
				}
			case stopAgents:
				agent.caseOfficers.Shutdown()
				agent.OnTrace(agent, "officers.Shutdown()")
			case startAgents:
				if agent.caseOfficers.Count() == 0 {
					initialize(agent, initAgent)
					agent.OnTrace(agent, "initialize()")
				}
			default:
				agent.Notify(common.MessageEventErrorStatus(agent.agentId, msg))
			}
		default:
		}
	}
}

func initialize(agent *ops, officer initOfficer) {
	if officer == nil {
		agent.Notify(core.NewStatusError(core.StatusInvalidArgument, errors.New("error: init officer is nil")))
		return
	}
	a := officer(westOrigin, agent)
	err := agent.caseOfficers.Register(a)
	if err != nil {
		agent.Notify(core.NewStatusError(core.StatusInvalidArgument, err))
	} else {
		a.Run()
	}
	a = officer(centralOrigin, agent)
	err = agent.caseOfficers.Register(a)
	if err != nil {
		agent.Notify(core.NewStatusError(core.StatusInvalidArgument, err))
	} else {
		a.Run()
	}
}
