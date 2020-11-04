package counter

import "testing"

func TestCounter(t *testing.T) {
	NotSafeCounter()
	MutexCounter()
	AtomicCounter()
	ChannelCounter()
}
