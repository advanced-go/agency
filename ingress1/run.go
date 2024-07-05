package ingress1

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

// run - case officer run function.
func run(c *controller) {
	if c == nil {
		return
	}
	tick := time.Tick(c.interval)

	for {
		select {
		case <-tick:
			status := core.StatusOK()
			if !status.OK() && !status.NotFound() {
				c.handler.Message(messaging.NewStatusMessage("", "", "", status))
			}
		case msg, open := <-c.ctrlC:
			if !open {
				return
			}
			switch msg.Event() {
			case messaging.ShutdownEvent:
				return
			default:
			}
		default:
		}
	}
}

/*
func updateAssignments(partition landscape1.Entry) ([]assignment1.Entry, *core.Status) {
	values := make(url.Values)
	values.Add(core.RegionKey, partition.Region)
	values.Add(core.ZoneKey, partition.Zone)
	values.Add(core.SubZoneKey, partition.SubZone)
	entries, _, status := assignment1.Get(nil, nil, values)
	return entries, status
}

func logActivity(body []activity1.Entry) *core.Status {
	_, status := activity1.Put(nil, body)
	return status
}

func processAssignments(c *caseOfficer, log func(body []activity1.Entry) *core.Status, update func(partition landscape1.Entry) ([]assignment1.Entry, *core.Status)) *core.Status {
	status := log([]activity1.Entry{{AgentId: c.uri}})
	if !status.OK() {
		return status
	}
	entries, status1 := update(c.partition)
	if !status1.OK() {
		return status
	}
	if len(entries) > 0 {
	}
	return status
}


*/
