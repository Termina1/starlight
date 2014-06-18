package main

import (
  "menteslibres.net/gosexy/redis"
)

type RedisConf struct {
  Host string
  Port uint
  Queue string
}

func CreateRedis(conf RedisConf) *redis.Client {
  r := redis.New()
  r.ConnectNonBlock(conf.Host, conf.Port)
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

func CreateSubscriptionRedis(conf RedisConf) <-chan string {
  r := CreateRedis(conf)
  return SubscribeRedis(r, conf.Queue)
}