package proxy

import "testing"

func TestExampleProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}
	sub.Proxy()

	// Output:
	// pre:real:after
}
