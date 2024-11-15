package service

import (
	"github.com/advanced-go/agency/common"
	"github.com/advanced-go/common/messaging"
)

// emissary attention
func emissaryAttend(r *ops) {
	for {
		// message processing
		select {
		case msg := <-r.emissary.C:
			switch msg.Event() {
			case messaging.ShutdownEvent:
				r.emissary.Close()
				return
			default:
				r.Handle(common.MessageEventErrorStatus(r.agentId, msg))
			}
		default:
		}
	}
}
