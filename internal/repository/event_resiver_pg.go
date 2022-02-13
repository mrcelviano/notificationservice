package repository

import "fmt"

const (
	logPrefix = "[DBR EVENT RECEIVER]"
)

type eventReceiver struct{}

func (er *eventReceiver) Event(eventName string) {
	fmt.Printf("%s called Event() event=%s\n", logPrefix, eventName)
}

func (er *eventReceiver) EventKv(eventName string, kvs map[string]string) {
	fmt.Printf("%s called EventKv() event=%s, kvs=%v\n", logPrefix, eventName, kvs)
}
func (er *eventReceiver) EventErr(eventName string, err error) error {
	fmt.Printf("%s called EventErr() event=%s, err=%v\n", logPrefix, eventName, err)
	return err
}
func (er *eventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	fmt.Printf("%s called EventErrKv() event=%s, err=%v, kvs=%v\n", logPrefix, eventName, err, kvs)
	return err
}
func (er *eventReceiver) Timing(eventName string, nanoseconds int64) {
	fmt.Printf("%s called Timing() event=%s nanoseconds=%d\n", logPrefix, eventName, nanoseconds)
}
func (er *eventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	fmt.Printf("%s called TimingKv() event=%s nanoseconds=%d, kvs=%v\n", logPrefix, eventName, nanoseconds, kvs)
}
