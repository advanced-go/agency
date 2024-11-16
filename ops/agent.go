package ops

import (
	"fmt"
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
	shutdownFunc func()
}

var opsAgent messaging.OpsAgent

func init() {
	opsAgent = NewAgent()
	opsAgent.Run()
}

// NewAgent - create a new ops agent
func NewAgent() messaging.OpsAgent {
	return newOpsAgent(Class)
}

func newOpsAgent(agentId string) *ops {
	r := new(ops)
	r.agentId = agentId
	r.caseOfficers = messaging.NewExchange()
	r.emissary = messaging.NewEnabledChannel()
	return r
}

// String - identity
func (o *ops) String() string { return o.Uri() }

// Uri - agent identifier
func (o *ops) Uri() string { return o.agentId }

func (o *ops) Handle(status *core.Status) *core.Status {
	var e core.Output
	return e.Handle(status)
}

func (o *ops) AddActivity(agentId string, content any) {
	fmt.Printf("%v : %v", agentId, content)
}

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

func shutdown(o *ops) {
	o.emissary.Close()
	o.caseOfficers.Shutdown()
}
