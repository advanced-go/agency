package ops

import (
	"errors"
	"github.com/advanced-go/common/core"
)

func initialize(agent *ops, newAgent newOfficerAgent) {
	if newAgent == nil {
		agent.Notify(core.NewStatusError(core.StatusInvalidArgument, errors.New("error: init officer is nil")))
		return
	}
	a := newAgent(westOrigin, agent)
	err := agent.caseOfficers.Register(a)
	if err != nil {
		agent.Notify(core.NewStatusError(core.StatusInvalidArgument, err))
	} else {
		a.Run()
	}
	a = newAgent(centralOrigin, agent)
	err = agent.caseOfficers.Register(a)
	if err != nil {
		agent.Notify(core.NewStatusError(core.StatusInvalidArgument, err))
	} else {
		a.Run()
	}
}
