package ops

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/common/test"
)

type statusT struct {
	status *core.Status
}

func (s *statusT) Notify(agent any, status *core.Status) *core.Status {
	s.status = status
	return status
}

func ExampleInitialize_Error() {
	notifier := new(statusT)
	agent := newOpsAgent(Class, notifier, test.Dispatcher)

	initialize(agent, nil)
	fmt.Printf("test: initialize() -> [status:%v]\n", notifier.status)

	notifier.status = nil
	initialize(agent, func(origin core.Origin, handler messaging.OpsAgent) messaging.OpsAgent {
		return test.NewAgent("", nil)
	})
	fmt.Printf("test: initialize() -> [status:%v]\n", notifier.status)

	notifier.status = nil
	a := test.NewAgent("agent:test", nil)
	err := agent.caseOfficers.Register(a)
	if err != nil {
		fmt.Printf("test: Register() -> [err:%v]\n", err)
	}
	initialize(agent, func(origin core.Origin, handler messaging.OpsAgent) messaging.OpsAgent {
		return a
	})
	fmt.Printf("test: initialize() -> [status:%v]\n", notifier.status)

	//Output:
	//test: initialize() -> [status:Invalid Argument [error: init officer is nil]]
	//test: initialize() -> [status:Invalid Argument [error: exchange.Register() agent Uri is nil]]
	//test: initialize() -> [status:Invalid Argument [error: exchange.Register() agent already exists: [agent:test]]]

}
