package ops

import (
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/log/timeseries"
	"github.com/advanced-go/resiliency/guidance"
)

const (
	PkgPath = "github/advanced-go/agency/ops"
)

var (
	westOrigin    = core.Origin{Region: "us-west", Host: "www.west-host1.com"}
	centralOrigin = core.Origin{Region: "us-central", Host: "www.central-host1.com"}
)

func StartAgents() {
	opsAgent.Message(messaging.NewControlMessage(opsAgent.Uri(), opsAgent.Uri(), startAgentsEvent))
	timeseries.Reset()
}

func StopAgents() {
	opsAgent.Message(messaging.NewControlMessage(opsAgent.Uri(), opsAgent.Uri(), stopAgentsEvent))
}

func SendCalendar() {
	msg := messaging.NewControlMessage(opsAgent.Uri(), opsAgent.Uri(), messaging.DataChangeEvent)
	msg.SetContent(guidance.ContentTypeCalendar, guidance.NewProcessingCalendar())
	opsAgent.Message(msg)
}
