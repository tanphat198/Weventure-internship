package GoLearning

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer,"Phat")

	got := buffer.String()
	want := "Hello, Phat"

	if got != want {
		t.Errorf("got '%s' but want '%s'",got ,want)
	}


}
