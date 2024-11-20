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
			agent.onMessage(msg, agent.emissary)
			switch msg.Event() {
			case messaging.ShutdownEvent:
				agent.finalize()
				return
			case messaging.DataChangeEvent:
				if msg.IsContentType(guidance.ContentTypeCalendar) {
					agent.caseOfficers.Broadcast(msg)
					agent.Trace(agent, "officers.Broadcast()")
				}
			case stopAgents:
				agent.caseOfficers.Shutdown()
				agent.Trace(agent, "officers.Shutdown()")
			case startAgents:
				if agent.caseOfficers.Count() == 0 {
					initialize(agent, initAgent)
					agent.Trace(agent, "initialize()")
				}
			default:
				agent.Notify(common.MessageEventErrorStatus(agent.agentId, msg))
			}
		default:
		}
	}
}
