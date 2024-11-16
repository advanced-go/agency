package common

import "github.com/advanced-go/common/messaging"

type Notification interface {
	OnTick(agent any)
	OnMessage(agent any, msg *messaging.Message)
}
