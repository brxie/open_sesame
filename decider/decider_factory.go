package decider

// NewDecider is factory to create new Decider.
func NewDecider(decider Decider, options ...interface{}) (Decider, error) {
	return decider.NewDecider(&options)
}