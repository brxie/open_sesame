package decider

// Decider is Decider Interface.
type Decider interface {
	Allowed() bool
	NewDecider(options *[]interface{}) (Decider, error)
}


