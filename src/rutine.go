package main

import (
  "fmt"
  "time"
)

func pinger(c chan string) {
  for i := 0; ; i++ {
    c <- "ping"
     time.Sleep(time.Second * 3)
  }
}

func ponger(c chan string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func printer(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string)

  go pinger(c)
  fmt.Println("test1")
  go ponger(c)
  fmt.Println("test2")
  go printer(c)
  fmt.Println("test3")


  var input string
  fmt.Scanln(&input)
}
