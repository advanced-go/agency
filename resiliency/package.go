package resiliency

import (
	"errors"
	"github.com/advanced-go/agency/ops"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/agency/resiliency"
	action  = "action"
	start   = "start"
	stop    = "stop"
	send    = "send"
)

// Post - resiliency POST
func Post(r *http.Request) (*http.Response, *core.Status) {
	if r == nil {
		status := core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
		return nil, status
	}
	switch r.URL.Query().Get("action") {
	case start:
		ops.StartAgents()
	case stop:
		ops.StopAgents()
	case send:
		ops.SendCalendar()
	}
	return httpx.NewResponse(http.StatusOK, nil, nil)
}
