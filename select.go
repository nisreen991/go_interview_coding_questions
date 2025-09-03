package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

/* In the program above, the server1 function in line no. 8 sleeps for 6 seconds then writes the text from server1 to the channel ch. The server2 function in line no. 12 sleeps for 3 seconds and then writes from server2 to the channel ch.

The main function calls the go Goroutines server1 and server2 in line nos. 20 and 21 respectively.

In line no. 22, the control reaches the select statement. The select statement blocks until one of its cases is ready. In our program above, the server1 Goroutine writes to the output1 channel after 6 seconds whereas the server2 writes to the output2 channel after 3 seconds. So the select statement will block for 3 seconds and will wait for server2 Goroutine to write to the output2 channel. After 3 seconds, the program prints, from server2 and terminates */

/*Practical Use of Select: Let’s assume we have a mission critical application and we need to return the output to the user as quickly as possible. The database for this application is replicated and stored in different servers across the world. Assume that the functions server1 and server2 are in fact communicating with 2 such servers. The response time of each server is dependant on the load on each and the network delay. We send the request to both the servers and then wait on the corresponding channels for the response using the select statement. The server which responds first is chosen by the select and the other response is ignored. This way we can send the same request to multiple servers and return the quickest response to the user */
