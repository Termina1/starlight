package main

import "github.com/golang/glog"

type StarBeam struct {
  MongoChannel <-chan string
  RedisChannel <-chan string
  Worker func(string)
  Density int
}


func CreateStarBeam(worker func(string), mconf MongoConf, rconf RedisConf, density int) *StarBeam {
  mongo := CreateSubscriptionMongo(mconf)
  glog.Info("Created subscription for MongoDB")
  redis := CreateSubscriptionRedis(rconf)
  glog.Info("Created subscription for Redis")
  return &StarBeam{mongo, redis, worker, density}
}

func (beam *StarBeam) launch() {
  out := beam.compose()
  glog.Info("Spawning ", beam.Density, " workers")
  for i := 0; i < beam.Density; i++ {
    go beam.spawn(out, i)
  }
}

func (beam *StarBeam) spawn(fan <-chan string, number int) {
  glog.Info("Worker #", number, " waiting")
  for repo := range fan {
    glog.Info("Worker #", number, " start processing", repo)
    beam.Worker(repo)
    glog.Info("Worker #", number, " stopped processing", repo)
  }
}

func (beam *StarBeam) compose() <-chan string {
  fanOut := make(chan string)
  go beam.prioritise(fanOut)
  return fanOut
}

func (beam *StarBeam) prioritise(fan chan string) {
  var repo string
  for {
    select {
      case repo = <- beam.RedisChannel:
      default:
        select {
          case repo = <- beam.RedisChannel:
          case repo = <- beam.MongoChannel:
        }
    }
    fan <- repo
  }
}
