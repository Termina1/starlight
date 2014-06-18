package main

import "github.com/stathat/jconfig"

func main() {
  config := jconfig.LoadConfig("config.json")
  mconf := MongoConf{
    config.GetString("localhost"),
    config.GetString("mongoDb"),
    config.GetString("mongoCollection"),
    config.GetInt("mongoBatch"),
    config.GetInt("mongoReindex"),
  }
  rconf := RedisConf{
    config.GetString("redisHost"),
    uint(config.GetInt("redisPort")),
    config.GetString("redisQueue"),
  }
  token := config.GetString("token")
  density := config.GetInt("beamDensity")
  beam := CreateStarBeam(StarExtractor(mconf, token), mconf, rconf, density)
  beam.launch()
  select {}
}