package main

import (
  "menteslibres.net/gosexy/redis"
)

func CreateRedis(host string, port uint) *redis.Client {
  r := redis.New()
  r.ConnectNonBlock(host, port)
  return r
}

func SubscribeRedis(client *redis.Client, queue string) <-chan string {
  tube := make(chan []string)
  go client.Subscribe(tube, queue)
  return createRedisFilter(tube)
}

func createRedisFilter(source chan []string)  <-chan string{
  out := make(chan string)
  go func() {
    for msg := range source {
      if (msg[0] == "message") {
        out <- msg[2]
      }
    }
  }()
  return out
}

func CreateSubscriptionRedis(host string, port uint, channel string) <-chan string {
  r := CreateRedis(host, port)
  return SubscribeRedis(r, channel)
}