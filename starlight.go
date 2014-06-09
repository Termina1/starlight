package main

import "fmt"

func main() {
  ch := CreateSubscriptionMongo("localhost", 100)
  fmt.Println(<-ch)
}