package GoLearning

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main(){
	Countdown(os.Stdout)
}

func Countdown(out io.Writer){
	for i := 3 ; i > 0 ; i--{
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out,i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out , "go")
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type Sleeper interface {
	Sleep()
}