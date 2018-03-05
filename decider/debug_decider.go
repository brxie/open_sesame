package decider

import (
	"math/rand"
	"time"
	"github.com/open_sesame/utils"
)

// DebugDecider is simple decider which
// will return configured allow decision.
type DebugDecider struct {
	config *DebugDeciderConfig
}

// NewDecider creates a new NewDecider.
func (_ DebugDecider) NewDecider(options *[]interface{}) (Decider, error) {
	// check if options are passed
	if len(*options) > 0 {
		config := (*options)[0].(DebugDeciderConfig)
		log.Log.Debug("Creating new decider with custom configuration")
		return DebugDecider{config: &config}, nil
	}

	dafaultConfig := DebugDeciderConfig{Mode: DENY}
	return DebugDecider{&dafaultConfig}, nil
}

// Allowed always returns true.
// May want use it for debug purposes.
func (a DebugDecider) Allowed() bool {
	switch mode := a.config.Mode; mode {
	case ALLOW:
		return true
	case DENY:
		return false
	case RANDOM:
		src := rand.NewSource(time.Now().UnixNano())
		randint := rand.New(src).Int()
		return randint % 2 == 0
	}
	return false
}

// DebugDeciderConfig is config used by DebugDecider
type DebugDeciderConfig struct {
	Mode mode
}

type mode int

const (
	// DENY always gives rejection decision
	DENY mode = iota
	// ALLOW always gives permit decision
	ALLOW
	// RANDOM gies random permit or reject decision
	RANDOM
)