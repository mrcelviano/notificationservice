package commons

import "fmt"

const (
	logPrefix = "[DBR EVENT RECEIVER]"
)

type dumbEventReceiver struct{}

func (er *dumbEventReceiver) Event(eventName string) {
	fmt.Printf("%s called Event() event=%s\n", logPrefix, eventName)
}

func (er *dumbEventReceiver) EventKv(eventName string, kvs map[string]string) {
	fmt.Printf("%s called EventKv() event=%s, kvs=%v\n", logPrefix, eventName, kvs)
}
func (er *dumbEventReceiver) EventErr(eventName string, err error) error {
	fmt.Printf("%s called EventErr() event=%s, err=%v\n", logPrefix, eventName, err)
	return err
}
func (er *dumbEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	fmt.Printf("%s called EventErrKv() event=%s, err=%v, kvs=%v\n", logPrefix, eventName, err, kvs)
	return err
}
func (er *dumbEventReceiver) Timing(eventName string, nanoseconds int64) {
	fmt.Printf("%s called Timing() event=%s nanoseconds=%d\n", logPrefix, eventName, nanoseconds)
}
func (er *dumbEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	fmt.Printf("%s called TimingKv() event=%s nanoseconds=%d, kvs=%v\n", logPrefix, eventName, nanoseconds, kvs)
}
