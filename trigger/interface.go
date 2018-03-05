
package trigger

type TriggerInterface interface {
	TriggersChan() chan bool
	RunAsync()
}


