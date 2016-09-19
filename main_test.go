package main
import "testing"
import "time"
import "fmt"

func TestSomething(t *testing.T) {
	fmt.Println(logo)
        fmt.Println("sleep 20 seconds")
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
		fmt.Printf("slept %v seconds\n", i)
	}
}
