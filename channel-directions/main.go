package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// | Function | Parameter | Type            | Direction    | Allowed Action |
// | -------- | --------- | --------------- | ------------ | -------------- |
// | `ping`   | `pings`   | `chan<- string` | send-only    | `pings <- msg` |
// | `pong`   | `pings`   | `<-chan string` | receive-only | `<-pings`      |
// | `pong`   | `pongs`   | `chan<- string` | send-only    | `pongs <- msg` |

// "passed message"
//       ↓
//   ping() → sends to → pings channel
//       ↓
//   pong() → receives from pings
//       ↓
//   pong() → sends to pongs
//       ↓
//   main() → receives from pongs
//       ↓
//   Output: passed message

// | Channel Direction | Syntax     | Meaning                                             |
// | ----------------- | ---------- | --------------------------------------------------- |
// | Send-only         | `chan<- T` | You can **only send** values of type `T` into it    |
// | Receive-only      | `<-chan T` | You can **only receive** values of type `T` from it |
