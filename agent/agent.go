package agent

import (
	"github.com/advanced-go/stdlib/messaging"
)

// All agents utilize: observations, guidance, and experience as input to inference -> to create actions
// observations + experience + guidance -> inference -> actions

// NewCaseOfficerAgent - assigned to a region/zone/sub-zone, managing unique hosts. Case officer agent manages
// host agents
func NewCaseOfficerAgent(uri string, assignment any, ctrlHandler messaging.Handler) (messaging.Agent, error) {
	return nil, nil
}

// NewHostAgent - agent responsible for a specific host in region/zone/sub-zone/host. Host agent manages
// ingress, egress traffic, and egress routing agents
func NewHostAgent(uri string, assignment any, ctrlHandler messaging.Handler) (messaging.Agent, error) {
	return nil, nil
}

// NewIngressAgent - ingress traffic controller for a host, one agent per host.
func NewIngressAgent(uri string, assignment any, ctrlHandler messaging.Handler) (messaging.Agent, error) {
	return nil, nil
}

// NewEgressTrafficAgent - egress traffic controller for a host, one agent per route
func NewEgressTrafficAgent(uri string, assignment any, ctrlHandler messaging.Handler) (messaging.Agent, error) {
	return nil, nil
}

// NewEgressRoutingAgent - egress traffic routing for a host, one agent per rout
func NewEgressRoutingAgent(uri string, assignment any, ctrlHandler messaging.Handler) (messaging.Agent, error) {
	return nil, nil
}
