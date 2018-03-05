
package trigger 

import (
	"time"
	"github.com/open_sesame/utils"
)

// PeriodicTrigger is simple trigger for periodic triggers generation.
// Helpful for debug purposes.
type PeriodicTrigger struct {
	triggersChan chan bool
}

const interval = 10 * time.Second

// NewPeriodicTrigger creates a new PeriodicTrigger.
func NewPeriodicTrigger(ch chan bool) *PeriodicTrigger {
	return &PeriodicTrigger{triggersChan: ch}
}

// TriggersChan returns the channel with triggered triggers.
func (b PeriodicTrigger) TriggersChan() chan bool {
	return b.triggersChan
}

// RunAsync Starts trigger in async mode.
func (b PeriodicTrigger) RunAsync() {
go func(c chan bool) {
	for {
		c <- true
		log.Log.Debug("Trigger generated!")
		time.Sleep(interval)
	}
	}(b.triggersChan)
}



