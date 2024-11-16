package ops

import (
	"github.com/advanced-go/agency/common"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/resiliency/guidance"
)

type newOfficer func(origin core.Origin, handler messaging.OpsAgent) messaging.OpsAgent

// emissary attention
func emissaryAttend(o *ops, newAgent newOfficer) {
	for {
		select {
		case msg := <-o.emissary.C:
			switch msg.Event() {
			case messaging.ShutdownEvent:
				o.shutdown()
				return
			case messaging.DataChangeEvent:
				if msg.IsContentType(guidance.ContentTypeCalendar) {
					o.caseOfficers.Broadcast(msg)
				}
			case stopAgents:
				o.caseOfficers.Shutdown()
			case startAgents:
				if o.caseOfficers.Count() == 0 {

				}
			default:
				o.Handle(common.MessageEventErrorStatus(o.agentId, msg))
			}
		default:
		}
	}
}
