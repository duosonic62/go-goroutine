package mainl

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuf bytes.Buffer
	defer stdoutBuf.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuf, "Producer Done.")
		for i := 0; i< 5; i++ {
			fmt.Fprintf(&stdoutBuf, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuf, "Received %v.\n", integer)
	}
}
