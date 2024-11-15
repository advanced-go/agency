package ops

import (
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/resiliency/guidance"
)

const (
	PkgPath = "github/advanced-go/agency/ops"
)

func StartAgents() {
	opsAgent.Message(messaging.NewControlMessage(opsAgent.Uri(), opsAgent.Uri(), startAgents))
}

func StopAgents() {
	opsAgent.Message(messaging.NewControlMessage(opsAgent.Uri(), opsAgent.Uri(), stopAgents))
}

func SendCalendar() {
	msg := messaging.NewControlMessage(opsAgent.Uri(), opsAgent.Uri(), messaging.DataChangeEvent)
	msg.SetContent(guidance.ContentTypeCalendar, guidance.NewProcessingCalendar())
	opsAgent.Message(msg)
}
