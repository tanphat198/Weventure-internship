package GoLearning

import "testing"

func TestHello(t *testing.T){

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' but want '%s'",got ,want)
		}
	}

	t.Run("in VietNam", func(t *testing.T) {
		got := Hello("Phat","VietNam")
		want := "Chao Phat"
		assertCorrectMessage(t, got, want)
	})
}