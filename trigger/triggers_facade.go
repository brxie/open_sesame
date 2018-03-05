package trigger

// Triggers is a facade used to initialize all triggers
// passed during initialization.
type Triggers struct {
	triggers []TriggerInterface
	triggersChan chan bool
}

// StartTriggers starts triggers in async mode.
func (t Triggers) StartTriggers() {
	for _, trig := range t.triggers {
		trig.RunAsync()
	}
}

// NewTriggers creates triggers
func NewTriggers() *Triggers{
	ch := make(chan bool)

	triggers := []TriggerInterface{&ButtonTrigger{triggersChan: ch},
								   &PeriodicTrigger{triggersChan: ch}}
	return &Triggers{triggers: triggers,
					 triggersChan: ch}
}

// TriggersChan returns chan which is passed to
// all created trigger instances.
func (t Triggers) TriggersChan() chan bool {
	return t.triggersChan
}