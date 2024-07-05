package ingress1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

const (
	Class = "ingress-controller1"
)

type ingressController struct {
	running  bool
	uri      string
	interval time.Duration
	ctrlC    chan *messaging.Message
	parent   messaging.Agent
	shutdown func()
}

func AgentUri(origin core.Origin) string {
	if origin.SubZone == "" {
		return fmt.Sprintf("%v:%v.%v.%v", Class, origin.Region, origin.Zone, origin.Host)
	}
	return fmt.Sprintf("%v:%v.%v.%v.%v", Class, origin.Region, origin.Zone, origin.SubZone, origin.Host)
}

// NewAgent - create a new case officer agent
func NewAgent(origin core.Origin, interval time.Duration, parent messaging.Agent) messaging.Agent {
	c := new(ingressController)
	c.uri = AgentUri(origin)
	c.interval = interval

	c.ctrlC = make(chan *messaging.Message, messaging.ChannelSize)
	c.parent = parent
	return c
}

// String - identity
func (c *ingressController) String() string {
	return c.uri
}

// Uri - agent identifier
func (c *ingressController) Uri() string {
	return c.uri
}

// Message - message the agent
func (c *ingressController) Message(m *messaging.Message) {
	messaging.Mux(m, c.ctrlC, nil, nil)
}

// Add - add a shutdown function
func (c *ingressController) Add(f func()) {
	c.shutdown = messaging.AddShutdown(c.shutdown, f)

}

// Shutdown - shutdown the agent
func (c *ingressController) Shutdown() {
	if !c.running {
		return
	}
	c.running = false
	if c.shutdown != nil {
		c.shutdown()
	}
	msg := messaging.NewControlMessage(c.uri, c.uri, messaging.ShutdownEvent)
	if c.ctrlC != nil {
		c.ctrlC <- msg
	}
}

// Run - run the agent
func (c *ingressController) Run() {
	if c.running {
		return
	}
	go run(c)
}
