
package trigger 

// ButtonTrigger is trigger for creating events on each push button.
type ButtonTrigger struct {
	triggersChan chan bool
}

// NewButtonTrigger creates a new NewButtonTrigger.
func NewButtonTrigger(ch chan bool) *ButtonTrigger {
	return &ButtonTrigger{triggersChan: ch}
}

// TriggersChan returns the channel with triggered triggers.
func (b ButtonTrigger) TriggersChan() chan bool {
	return b.triggersChan
}

// RunAsync Starts trigger in async mode.
func (b ButtonTrigger) RunAsync() {
go func(c chan bool) {
		// TODO: GPIO input implementation
	}(b.triggersChan)
}

