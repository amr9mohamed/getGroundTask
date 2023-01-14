package guests

import (
	"time"
)

// todo :: add checked_out flag

type Guest struct {
	Name         string
	TableID      uint
	Accompanying int64
	TimeArrived  *time.Time
}
