package guests

import (
	"time"
)

type Guest struct {
	Name         string
	TableID      uint
	Accompanying int64
	TimeArrived  time.Time
}
