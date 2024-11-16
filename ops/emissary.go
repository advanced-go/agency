package ops

import (
	"github.com/advanced-go/agency/common"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/resiliency/guidance"
)

type initOfficer func(origin core.Origin, handler messaging.OpsAgent) messaging.OpsAgent

// emissary attention
func emissaryAttend(o *ops, agent initOfficer) {
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
					initialize(o, agent)
				}
			default:
				o.Handle(common.MessageEventErrorStatus(o.agentId, msg))
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
}
