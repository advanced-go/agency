package ops

import (
	"github.com/advanced-go/agents/caseofficer"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
)

const (
	Class = "agency-ops"
)

type ops struct {
	running bool
	agentId string

	emissary     *messaging.Channel
	caseOfficers *messaging.Exchange
	notifier     messaging.Notifier
	tracer       messaging.Tracer
	sender       dispatcher
}

func cast(agent any) *ops {
	o, _ := agent.(*ops)
	return o
}

var opsAgent messaging.OpsAgent

func init() {
	opsAgent = NewAgent()
	opsAgent.Run()
}

// NewAgent - create a new ops agent
func NewAgent() messaging.OpsAgent {
	return newOpsAgent(Class, messaging.LogErrorNotifier, messaging.DefaultTracer, newDispatcher())
}

func newOpsAgent(agentId string, notifier messaging.Notifier, tracer messaging.Tracer, sender dispatcher) *ops {
	r := new(ops)
	r.agentId = agentId
	r.caseOfficers = messaging.NewExchange()
	r.emissary = messaging.NewEmissaryChannel(true)
	r.notifier = notifier
	r.tracer = tracer
	r.sender = sender
	return r
}

// String - identity
func (o *ops) String() string { return o.Uri() }

// Uri - agent identifier
func (o *ops) Uri() string { return o.agentId }

// Notify - status notifier
func (o *ops) Notify(status *core.Status) *core.Status { return o.notifier.Notify(status) }

// Trace - activity tracing
func (o *ops) Trace(agent messaging.Agent, event, activity string) {
	o.tracer.Trace(agent, event, activity)
}

// Message - message the agent
func (o *ops) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	o.emissary.C <- m
}

// Run - run the agent
func (o *ops) Run() {
	if o.running {
		return
	}
	go emissaryAttend(o, caseofficer.NewAgent)
	o.running = true
}

// Shutdown - shutdown the agent
func (o *ops) Shutdown() {
	if !o.running {
		return
	}
	o.running = false
	msg := messaging.NewControlMessage(o.Uri(), o.Uri(), messaging.ShutdownEvent)
	o.emissary.Enable()
	o.emissary.C <- msg
}

func (o *ops) IsFinalized() bool {
	return o.emissary.IsFinalized() && o.caseOfficers.IsFinalized()
}

func (o *ops) finalize() {
	o.emissary.Close()
	o.caseOfficers.Shutdown()
}

func (o *ops) setup(event string) {
	o.sender.setup(o, event)
}

func (o *ops) dispatch(event string) {
	o.sender.dispatch(o, event)
}
