package egress1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

const (
	Class = "egress-controller1"
)

type controller struct {
	running  bool
	uri      string
	interval time.Duration // Needs to be configured dynamically during runtime
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
func NewAgent(origin core.Origin, parent messaging.Agent) messaging.Agent {
	c := new(controller)
	c.uri = AgentUri(origin)
	//c.interval = interval

	c.ctrlC = make(chan *messaging.Message, messaging.ChannelSize)
	c.parent = parent
	return c
}

// String - identity
func (c *controller) String() string {
	return c.uri
}

// Uri - agent identifier
func (c *controller) Uri() string {
	return c.uri
}

// Message - message the agent
func (c *controller) Message(m *messaging.Message) {
	messaging.Mux(m, c.ctrlC, nil, nil)
}

// Add - add a shutdown function
func (c *controller) Add(f func()) {
	c.shutdown = messaging.AddShutdown(c.shutdown, f)

}

// Shutdown - shutdown the agent
func (c *controller) Shutdown() {
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
func (c *controller) Run() {
	if c.running {
		return
	}
	go run(c)
}
