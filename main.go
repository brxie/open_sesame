package main

import (
	"github.com/open_sesame/trigger"
	"github.com/open_sesame/decider"
	"github.com/open_sesame/utils"
)


func main() {
	log.Log.Debug("Open Sesame start!")
	
	// triggers
	triggers := trigger.NewTriggers()
	triggers.StartTriggers()	

	// deciders
	deciders := []decider.Decider{}
	config := decider.LicencePlatesDeciderConfig{
		Plates: []string{"WF65464", "KR3F257"},
	}
	platesDecider, _ := decider.NewDecider(decider.LicencePlatesDecider{}, config)
	deciders = append(deciders, platesDecider)
	log.Log.Infof("%d deciders initialized sucessfully!", len(deciders))

	// main
	for {
		<-triggers.TriggersChan()
		for _, decider := range deciders {
			if decider.Allowed() == true {
				// do some magic!
				break
			}
		}
	}
}
