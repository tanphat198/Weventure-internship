package GoLearning

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool{
	ressults := make(map[string]bool)

	for _, url := range urls {
		ressults[url] = wc(url)
	}
	return ressults
}

////////////////////////////////////////////////////////////////////

func mockWebsiteChecker(url string) bool{
	if url == "waat://furhurterwe.geds"{
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":			true,
		"http://blog.gypsydave5.com":	true,
		"waat://furhurterwe.geds":		false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got){
		t.Fatalf("wanted %v but got %v",want ,got)
	}
}


func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
//func main() {
//	array := []int{}
//	chunk1 := 0
//	chunk2 := 0
//	done1 := make(chan bool)
//	done2 := make(chan bool)
//	for i := 0; i < 20; i++ {
//		array = append(array, i)
//	}
//	fmt.Printf("%v\n", array)
//	go func(done1 chan bool) {
//		for i := 0; i < 10; i++ {
//			chunk1 += array[i]
//		}
//		fmt.Printf("chunk1= %v\n", chunk1)
//		done1 <- true
//	}(done1)
//	go func(done2 chan bool) {
//		for i := 10; i < 20; i++ {
//			chunk2 += array[i]
//		}
//		fmt.Printf("chunk2= %v\n", chunk2)
//		done2 <- true
//	}(done2)
//	<-done1
//	<-done2
//	fmt.Printf("chunk1 + chunk2= %v\n", chunk1+chunk2)
//	time.Sleep(2 * time.Second)
//}
