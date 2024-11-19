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
	running      bool
	agentId      string
	emissary     *messaging.Channel
	caseOfficers *messaging.Exchange
	tracer       messaging.Tracer
	notifier     messaging.Notifier
	dispatcher   messaging.Dispatcher
	shutdownFunc func()
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
	return newOpsAgent(Class, messaging.DefaultTracer, messaging.LogErrorNotifier, messaging.MutedDispatcher)
}

func newOpsAgent(agentId string, tracer messaging.Tracer, notifier messaging.Notifier, dispatcher messaging.Dispatcher) *ops {
	r := new(ops)
	r.agentId = agentId
	r.caseOfficers = messaging.NewExchange()
	r.emissary = messaging.NewEmissaryChannel(true)
	r.tracer = tracer
	r.notifier = notifier
	r.dispatcher = dispatcher
	return r
}

// String - identity
func (o *ops) String() string { return o.Uri() }

// Uri - agent identifier
func (o *ops) Uri() string { return o.agentId }

// Trace - agent activity tracing
func (o *ops) Trace(agent messaging.Agent, activity any) {
	o.tracer.Trace(agent, activity)
}

// Notify - status notifier
func (o *ops) Notify(status *core.Status) *core.Status {
	return o.notifier.Notify(status)
}

func (o *ops) OnTick(agent any, src *messaging.Ticker) { o.dispatcher.OnTick(agent, src) }
func (o *ops) OnMessage(agent any, msg *messaging.Message, src *messaging.Channel) {
	o.dispatcher.OnMessage(agent, msg, src)
}
func (o *ops) OnError(agent any, status *core.Status) *core.Status { return o.OnError(agent, status) }

// Message - message the agent
func (o *ops) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	o.emissary.C <- m
}

// Add - add a shutdown function
func (o *ops) Add(f func()) { o.shutdownFunc = messaging.AddShutdown(o.shutdownFunc, f) }

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
	if o.shutdownFunc != nil {
		o.shutdownFunc()
	}
	msg := messaging.NewControlMessage(o.agentId, o.agentId, messaging.ShutdownEvent)
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
