package ops

import (
	"errors"
	"github.com/advanced-go/common/core"
)

func initialize(agent *ops, officer initOfficer) {
	if officer == nil {
		agent.Notify(agent, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: init officer is nil")))
		return
	}
	a := officer(westOrigin, agent)
	err := agent.caseOfficers.Register(a)
	if err != nil {
		agent.Notify(agent, core.NewStatusError(core.StatusInvalidArgument, err))
	} else {
		a.Run()
	}
	a = officer(centralOrigin, agent)
	err = agent.caseOfficers.Register(a)
	if err != nil {
		agent.Notify(agent, core.NewStatusError(core.StatusInvalidArgument, err))
	} else {
		a.Run()
	}
}
