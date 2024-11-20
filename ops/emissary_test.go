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
	agent := newOpsAgent(Class, messaging.OutputErrorNotifier, test.DefaultTracer)

	go func() {
		go emissaryAttend(agent, nil)
		agent.Message(shutdownMsg)
		fmt.Printf("test: Shutdown() -> [finalized:%v]\n", agent.IsFinalized())
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//OnMsg()   -> agency-ops : event:shutdown channel:EMISSARY
	//test: Shutdown() -> [finalized:true]

}

func ExampleEmissary_Stop() {
	ch := make(chan struct{})
	agent := newOpsAgent(Class, messaging.OutputErrorNotifier, test.DefaultTracer)

	go func() {
		go emissaryAttend(agent, nil)
		agent.Message(stopMsg)
		agent.Message(shutdownMsg)
		agent.IsFinalized()
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//OnMsg()   -> agency-ops : event:stop-agents channel:EMISSARY
	//OnTrace() -> agency-ops : officers.Shutdown()
	//OnMsg()   -> agency-ops : event:shutdown channel:EMISSARY

}

func ExampleEmissary_DataChange() {
	ch := make(chan struct{})
	agent := newOpsAgent(Class, messaging.OutputErrorNotifier, test.DefaultTracer)

	go func() {
		go emissaryAttend(agent, nil)
		agent.Message(dataChangeMsg)
		agent.Message(shutdownMsg)
		agent.IsFinalized()
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//OnMsg()   -> agency-ops : event:data-change channel:EMISSARY
	//OnTrace() -> agency-ops : officers.Broadcast()
	//OnMsg()   -> agency-ops : event:shutdown channel:EMISSARY

}

func ExampleEmissary_Start_Error() {
	ch := make(chan struct{})
	agent := newOpsAgent(Class, messaging.OutputErrorNotifier, test.DefaultTracer)

	go func() {
		go emissaryAttend(agent, nil)
		agent.Message(startMsg)
		agent.Message(shutdownMsg)
		agent.IsFinalized()
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//OnMsg()   -> agency-ops : event:start-agents channel:EMISSARY
	//{ "timestamp":"2024-11-19T21:36:19.968Z", "code":3, "status":"Invalid Argument", "request-id":null, "errors" : [ "error: init officer is nil" ], "trace" : [ "https://github.com/advanced-go/common/tree/main/messaging.(*outputError)#Notify","https://github.com/advanced-go/agency/tree/main/ops#initialize" ] }
	//OnTrace() -> agency-ops : initialize()
	//OnMsg()   -> agency-ops : event:shutdown channel:EMISSARY

}

func ExampleEmissary_Start() {
	ch := make(chan struct{})
	agent := newOpsAgent(Class, messaging.OutputErrorNotifier, test.DefaultTracer)

	go func() {
		go emissaryAttend(agent, nil)
		agent.Message(startMsg)
		agent.Message(shutdownMsg)
		agent.IsFinalized()
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//{ "timestamp":"2024-11-19T21:36:19.968Z", "code":3, "status":"Invalid Argument", "request-id":null, "errors" : [ "error: init officer is nil" ], "trace" : [ "https://github.com/advanced-go/common/tree/main/messaging.(*outputError)#Notify","https://github.com/advanced-go/agency/tree/main/ops#initialize" ] }
	//OnError() -> agency-ops : Invalid Argument [error: init officer is nil]
	//OnTrace() -> agency-ops : initialize()
	//OnMsg()   -> agency-ops : event:start-agents channel:EMISSARY
	//OnMsg()   -> agency-ops : event:shutdown channel:EMISSARY

}
