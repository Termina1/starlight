package main

type StarBeam struct {
  MongoChannel <-chan string
  RedisChannel <-chan string
  Worker func(string)
  Density int
}


func CreateStarBeam(worker func(string), mconf MongoConf, rconf RedisConf, density int) *StarBeam {
  mongo := CreateSubscriptionMongo(mconf)
  redis := CreateSubscriptionRedis(rconf)
  return &StarBeam{mongo, redis, worker, density}
}

func (beam *StarBeam) launch() {
  out := beam.compose()
  for i := 0; i < beam.Density; i++ {
    go beam.spawn(out)
  }
}

func (beam *StarBeam) spawn(fan <-chan string) {
  for repo := range fan {
    beam.Worker(repo)
  }
}

func (beam *StarBeam) compose() <- chan string {
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