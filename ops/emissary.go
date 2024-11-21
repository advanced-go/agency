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
	agent.dispatch(messaging.StartupEvent)
	for {
		select {
		case msg := <-agent.emissary.C:
			agent.setup(msg.Event())
			switch msg.Event() {
			case messaging.ShutdownEvent:
				agent.finalize()
				agent.dispatch(msg.Event())
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
					initialize(agent, initAgent)
					agent.dispatch(msg.Event())
				}
			default:
				agent.Notify(common.MessageEventErrorStatus(agent.agentId, msg))
			}
		default:
		}
	}
}
