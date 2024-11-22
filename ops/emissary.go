package ops

import (
	"github.com/advanced-go/agency/common"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/resiliency/guidance"
)

type newOfficerAgent func(origin core.Origin, handler messaging.OpsAgent) messaging.OpsAgent

// emissary attention
func emissaryAttend(agent *ops, newAgent newOfficerAgent) {
	// Agent is always running
	//agent.dispatch(messaging.StartupEvent)
	//agent.dispatch(msg.Event())
	for {
		select {
		case msg := <-agent.emissary.C:
			agent.setup(msg.Event())
			switch msg.Event() {
			case messaging.ShutdownEvent:
				agent.finalize()
				return
			case messaging.DataChangeEvent:
				if msg.IsContentType(guidance.ContentTypeCalendar) {
					agent.caseOfficers.Broadcast(msg)
					agent.dispatch(msg.Event())
				}
			case stopAgentsEvent:
				agent.caseOfficers.Shutdown()
				agent.dispatch(msg.Event())
			case startAgentsEvent:
				if agent.caseOfficers.Count() == 0 {
					initialize(agent, newAgent)
					agent.dispatch(msg.Event())
				}
			default:
				agent.Notify(common.MessageEventErrorStatus(agent.agentId, msg))
			}
		default:
		}
	}
}
