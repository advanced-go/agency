package ops

import (
	"fmt"
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/common/test"
	"github.com/advanced-go/resiliency/guidance"
)

var (
	shutdownMsg   = messaging.NewControlMessage("", "", messaging.ShutdownEvent)
	dataChangeMsg = messaging.NewControlMessage("", "", messaging.DataChangeEvent)
	startMsg      = messaging.NewControlMessage("", "", startAgents)
	stopMsg       = messaging.NewControlMessage("", "", stopAgents)
)

func init() {
	dataChangeMsg.SetContent(guidance.ContentTypeCalendar, guidance.NewProcessingCalendar())
}

func ExampleEmissary_Shutdown() {
	ch := make(chan struct{})
	agent := newOpsAgent(Class, messaging.DefaultTracer, messaging.OutputErrorNotifier, test.Dispatcher)

	go func() {
		go emissaryAttend(agent, nil)
		agent.Message(shutdownMsg)
		fmt.Printf("\ntest: Shutdown() -> [finalized:%v]\n", agent.IsFinalized())
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//OnMsg() -> agency-ops : event:shutdown channel:EMISSARY
	//test: Shutdown() -> [finalized:true]

}
